package sprint

import (
	"fmt"
	"strconv"
	"testing"
)

type F float64

func (f F) String() string {
	return strconv.FormatFloat(float64(f), 'f', 6, 64)
}

var f = F(1.234)

func TestSprint(t *testing.T) {
	testset := []interface{}{
		f,
		"this is string",
		int(100),
		1 == 1,
		1 == 2,
		float64(1.2),
	}

	for _, v := range testset {
		fmt.Println(Sprint(v))
	}
}
