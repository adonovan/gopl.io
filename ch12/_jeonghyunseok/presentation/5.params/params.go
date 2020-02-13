// tag 를 이용한 파싱으로 보인다

package params

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i:=0; i<v.NumField(); i++{
		fieldInfo:= v.Type().Field(i)
		tag:=fieldInfo.tag
		name:=tag.Get("http")
		if name ==""{
			name = strings.ToLower(fieleInfo.Name)
		}
		fields[name] = v.Field(i)
	}


	for name, values := range req.Form{
		f:= fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice{
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))

			} else {
				if err := populate(f, value); err !=nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}


func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err !=nil {
			return err
		}
		v.SetInt(i)

	case reflectBool:
		b, err := strconv.ParseBool(value)
		if err!=nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind: %s", v.Type())
	}
	return nil
	
}