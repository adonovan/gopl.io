package main

import (
	"fmt"
	. "net"
)

func isUp(v Flags) bool     { return v&FlagUp == FlagUp }
func turnDown(v *Flags)     { *v &^= FlagUp }
func setBroadcast(v *Flags) { *v |= FlagBroadcast }
func isCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t \n", v, isUp(v))
	turnDown(&v)
	fmt.Printf("%b %t \n", v, isUp(v))
	setBroadcast(&v)
	fmt.Printf("%b %t \n", v, isUp(v))
	fmt.Printf("%b %t \n", v, isCast(v))
}
