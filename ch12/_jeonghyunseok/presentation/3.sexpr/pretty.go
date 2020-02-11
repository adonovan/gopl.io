// 이쁘게 프린팅 하는 법. 역사가 깊네요.

package sexpr

import (
	"reflect"
	"bytes"
	"fmt"
)

const margin = 80

func MarshalIndent(v interface{}) ([]byte, error) {
	p:=printer{width:margin}
	if err:= pretty(&p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

type token struct {
	kind rune
	str string
	size int
}

type printer struct{
	tokens []*token
	stack []*token
	rtotal int

	bytes.Buffer
	indent []int
	width int
}

func (p *printer) string(str string) {
	tok := &token{
		kind: 's',
		str: str,
		size:len(str),
	}
	if len(p.stack)==0{
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotal +=len(str)
	}
}

func (p *printer) pop() (tok *token) {
	last := len(p.stack) -1
	tok, p.stack = p.stack[last], p.stack[:last]
	return
}

func (p *printer) begin() {
	if len(p.stack) == 0 {
		p.rtotal = 1
	}
	t:= &token{
		kind: '(',
		size: -p.rtotal}
		p.tokens = append(p.tokens, t)
		p.stack = append(p.stack, t)
		p.string("(")
}

func (p *printer) end() {
	p.string(")")
	p.tokens = append(p.tokens, &token{kind: ')'})
	x := p.pop()
	x.size += p.rtotal
	if x.kind == ' ' {
		p.pop().size += p.rtotal
	}
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}


func (p *printer) space() {
	last := len(p.stack) -1
	x:= p.stack[last]
	if x.kind == ' ' {
		x.size +=p.rtotal
		p.stack = p.stack[:last] 
	}
	t:= &token{
		kind: ' ',
		size: -p.rtotal,
	}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}

func (p *printer) print(t *token) {
	switch t.kind{
	case 's':
		p.WriteString(t.str)
		p.width -=len(t.str)
	case '(':
		p.indent = append(p.indent, p.width )
	case ')':
		p.indent = p.indent[:len(p.indent)-1]
	case ' ':
		if t.size>p.width{
			p.width = p.indent[len(p.indent)-1] -1
			fmt.Fprintf(&p.Buffer, "\n%*s", margin-p.width, "")
		} else {
			p.WriteByte(' ')
			p.width--
		}
	}
}

func (p *printer) stringf(format string, args ...interface{}){
	p.string(fmt.Sprintf(format, args...))
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("nil")
		case reflect.Int, reflect.Int8, reflect.Int16, 
			reflect.Int32, reflect.Int64:
			p.stringf("%d", v.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, 
			reflect.Uint32, reflect.Uint64:
			p.stringf("%d", v.Uint())

	case reflect.String:
		p.stringf("%q", v.String())
	case reflect.Array, reflect.Slice:
		p.begin()
		for i:=0; i<v.Len(); i++{
			if i>0{
				p.space()
			}
			if err:= pretty(p, v.Index(i)); err != nil {
				return err 
			}
		}
		p.end()
	case reflect.Struct:
		p.begin()
		for i:=0; i<v.NumField(); i++ {
			if i>0 {
				p.space()
			}
			p.begin()
			p.string(v.Type().Field(i).Name)
			p.space()
			if err:= pretty(p, v.Field(i)); err!= nil{
				return err
			}
			
			p.end()
		}
			p.end()
	case reflect.Map:
		p.begin()
		for i, key := range v.MapKeys() {
			if i > 0{
				p.space()
			}
			p.begin()
			if err := pretty(p, key); err != nil {
				return err
			}
			p.space()
			if err:= pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
			p.end()
		}
		p.end()
	case reflect.Ptr:
		return pretty(p, v.Elem())

	default: 
	return fmt.Errorf("unsupperted type: %s", v.Type())
	}
	return nil
}