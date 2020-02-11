// 이쁘게 프린팅 하는 법. 역사가 깊네요.

package sexpr

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
		size:len(str)}
	}
	if len(p.stack)==0{
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotoal +=len(str)
	}
}

func (p *printer) pop() (tok *token) {
	last := len(p.stack) -1
	top, p.stack = p.stack[last], p.stack[:last]
	return
}

