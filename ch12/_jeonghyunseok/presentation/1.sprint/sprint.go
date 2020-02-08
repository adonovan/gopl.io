//

package sprint

import (
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {

	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x == true {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}
