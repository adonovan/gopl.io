package format

import (
	"reflect"
	"strconv"
)
// get any interface and return string
// SEE it's hiding reflect.ValueOf() function - not exposing
func Any(i interface{}) string {
	return formatAtom(reflect.ValueOf(i))
}


func formatAtom(v reflect.Value) string {
	switch v.Kind() { // this example code is about Kind() func
	case reflect.Invalid:
		return "invalid"
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			return strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}

}