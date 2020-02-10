// 메쏘드는 뭐하는 녀석일까? 
// 어떠한 value 이든 그 method 를 출력해준다.

package methods

import (
	"fmt"
	"reflect"
	"strings"
)


func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i:=0; i<v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func(%s) %s%s\n", t, t.Method(i).Name, 
			strings.TrimPrefix(methType.String(), "func"))
	}
}