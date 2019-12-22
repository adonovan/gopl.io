// Netflag 는 정수 타입을 bit 필드로 사용하는 것을 보여준다.
package main

import (
	"fmt"
	"net"
)

// type Flags uint

// const (
// 	FlagUp           Flags = 1 << iota // 0x00001 is up
// 	FlagBroadcast                      // 0x00010 supports broadcast access capability
// 	FlagLoopback                       // 0x00100 is a loopback interface
// 	FlagPointToPoint                   // 0x01000 belongs to a pointtopoint // 	link
// 	FlagMulticast                      // 0x10000 supports multicast access capability
// )

func IsUp(v net.Flags) bool     { return v&net.FlagUp == net.FlagUp }
func TurnDown(v *net.Flags)     { *v &^= net.FlagUp }
func SetBroadcast(v *net.Flags) { *v |= net.FlagBroadcast }
func IsCast(v net.Flags) bool   { return v&(net.FlagBroadcast|net.FlagMulticast) != 0 }

func main() {
	var v net.Flags = net.FlagMulticast | net.FlagUp // 0x10001
	// net.Flags 타입의 v 는 Multicast 와 Up 두 개의 정보를 담고 있다.

	fmt.Printf("%b, %t\n", v, IsUp(v)) // 10001 true
	// v 가 담고 있는 FlagUp 이 Up 인가 확인

	TurnDown(&v)
	// 이건 좀 헷갈리기도 하는데
	// 일단 FlagUp 비트를 0으로 만드는 것인가 보다
	// 순서대로 &= 해주고 그 결과를 ^= 해주는 것. bit clear 연산자이다.

	fmt.Printf("%b, %t\n", v, IsUp(v)) // 10000 false

	SetBroadcast(&v)
	fmt.Printf("%b, %t\n", v, IsUp(v))   // 10010 false
	fmt.Printf("%b, %t\n", v, IsCast(v)) // 10010 true

}

/*

go run main.go

*/
