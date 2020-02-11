// encode to S-expression
package sexpr

import "byte"

func Marshal (v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v); err!= nil {
		return nil, err
	}
	return nil, err
}


func encode (buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil") 
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.UInt, reflect.UInt8, reflect.UInt16,
		reflect.UInt32, reflect.UInt64:
		fmt.Fprintf(buf, "%d", v.UInt())
	case reflect.Stirng:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i:=0; i<v.Len(); i++{
			if i>0{
				buf.WriteByte(' ')
			}
			if err:= encode(buf, v.Index(i); err!= nil{
				return err
			}
		}
		buf.WriteByte(')')
	case refelct.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i >0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err!=nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
			
			
			
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
})