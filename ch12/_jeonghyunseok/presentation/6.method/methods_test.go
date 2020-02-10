// 실제 테스트를 보자

package methods_test

import (
	"strings"
	"time"

	"github.com/gopl-study/gopl.io/ch12/_jeonghyunseok/presentation/methods"
)

func ExamplePrintDuration() {
	methods.Print(time.Hour)
	// Output:
	// type time.Duration
	// func(time.Duration) Hours() float64
	// func(time.Duration) Microseconds() int64
	// func(time.Duration) Milliseconds() int64
	// func(time.Duration) Minutes() float64
	// func(time.Duration) Nanoseconds() int64
	// func(time.Duration) Round(time.Duration) time.Duration
	// func(time.Duration) Seconds() float64
	// func(time.Duration) String() string
	// func(time.Duration) Truncate(time.Duration) time.Duration
}

/*
아래와 같이 지원 메쏘드가 늘어났다.
got:
type time.Duration
func(time.Duration) Hours() float64
func(time.Duration) Microseconds() int64
func(time.Duration) Milliseconds() int64
func(time.Duration) Minutes() float64
func(time.Duration) Nanoseconds() int64
func(time.Duration) Round(time.Duration) time.Duration
func(time.Duration) Seconds() float64
func(time.Duration) String() string
func(time.Duration) Truncate(time.Duration) time.Duration
want:
type time.Duration
func (time.Duration) Hours() float64
func (time.Duration) Minutes() float64
func (time.Duration) Nanoseconds() int64
func (time.Duration) Seconds() float64
func (time.Duration) String() string
--- FAIL: ExamplePrintReplacer (0.00s)
*/

func ExamplePrintReplacer() {
	methods.Print(new(strings.Replacer))
	// Output:
	// type *strings.Replacer
	// func(*strings.Replacer) Replace(string) string
	// func(*strings.Replacer) WriteString(io.Writer, string) (int, error)
}

/*

이건 이전하고 똑같으니 성공할줄 알았는데 띄어쓰기 틀려서 실패. 허허
got
type *strings.Replacer
func(*strings.Replacer) Replace(string) string
func(*strings.Replacer) WriteString(io.Writer, string) (int, error)
want:
type *strings.Replacer
func (*strings.Replacer) Replace(string) string
func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
*/
