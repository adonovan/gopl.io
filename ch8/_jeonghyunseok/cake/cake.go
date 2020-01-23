// 케이크는 다양한 파라미터를 가진 동시성을 가진 케이크 가게 시뮬이다.package cake
// go test -bench=. gopl.io/ch8/cake 라고 하면 벤치마크 가능하다.
// 여기서 .은 모든 테스트를 의미한다. 정규식이다.

package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool          // 떠벌릴 것인가?
	Cakes          int           // 구워야 할 빵 개수
	BakeTime       time.Duration // 빵하나 굽는 시간
	BakeStdDev     time.Duration // 빵하나 굽는데 대한 표준 편차 http://bit.ly/38vAHNZ
	BakeBuf        int           // 빵 굽고 식히는 사이의 버퍼
	NumIcers       int           // 식히는 빵 개수
	IceTime        time.Duration // 빵하나 식히는 시간
	IceStdDev      time.Duration // 빵하나 식히는 표준편차
	IceBuf         int           // 식히고 장식 쓰는데 걸리는 시간
	InscribeTime   time.Duration // 빵하나 장식 시간
	InscribeStdDev time.Duration // 장식 글쓰는 표준편차
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

/*
분석해보자
1. 총 s.Cake 개의 빵을 구워야 한다.
2. 각 빵은 0 부터 s.Cake-1 까지의 번호를 부여받는다. Verbose 상태면 몇 번빵 굽는지 프린트
3. work() 를 한 번 실행한다. 이건 빵 굽는 시간을 평균 + 표준편차를 감안해서 그냥 delay 주는 함수이다.
4. 그러고 나서는 baked 라는 (아마도 buffered) 채널에 써준다. "빵 다 구웠어요"
5. 결국 baked 채널에는 빵이 구워지는 대로 들어가게 된다.

*/

func (s *Shop) icer(icerNum int, iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icerNum:", icerNum, ", icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(i, iced, baked)
		}
		s.inscriber(iced)
	}
}

/*
Baker 와 Inscriber 는 하나씩 뿐이지만
Icer 는 여럿이다.
그래서 Icer 번호도 표시하게 바꿔보았다.
*/

// 대략 빵 하나 굽는 동안 슬립 들어가는 함수
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}
