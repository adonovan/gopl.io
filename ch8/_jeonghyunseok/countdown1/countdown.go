// 카운트다운은 로켓 발사 카운트 다운을 구현한것이다.package countdown1

package main

import (
	"log"
	"fmt"
	"time"
)

func main() {
	log.Print("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for cd := 10; cd > 0; cd-- {
		fmt.Print(cd)
		fmt.Println(": ", <-tick)
	}
	launch()
}

func launch() {
	log.Print("Lift off!")
}
