package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	nums := make([]float64, 0, 0)
	if len(os.Args) < 2 {
		fmt.Print("Enter text:")

		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		n, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, n)
	} else {
		for _, arg := range os.Args[1:] {
			n, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			nums = append(nums, n)
		}
	}

	for _, num := range nums {
		kg := KG(num)
		fmt.Println(kg.String()+":", KGtoLB(kg))

		lb := LB(num)
		fmt.Println(lb.String()+":", LBtoKG(lb))
	}
}

// LB 파운드
type LB float64

// KG 킬로그램
type KG float64

func (l LB) String() string {
	return fmt.Sprintf("%glb", l)
}

func (k KG) String() string {
	return fmt.Sprintf("%gkg", k)
}

// KGtoLB kg to lb
func KGtoLB(k KG) LB {
	return LB(k) * 2.205
}

// LBtoKG lb to kg
func LBtoKG(l LB) KG {
	return KG(l) / 2.205
}
