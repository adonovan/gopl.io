package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"gopl.io/ch8/thumbnail"
)

func main() {
	start := time.Now()
	imgs := make([]string, 0)
	for i := 0; i < 200; i++ {
		for j := 0; j < 5; j++ {
			num := strconv.Itoa(j)
			imgs = append(imgs, "img"+num+".jpg")
		}
	}
	makeThumnails(imgs)
	// makeThumnails3(imgs)
	fmt.Println(time.Since(start))
}

// makeThumnails 작업, 아래에서 고루틴을 이용하여 개선.
func makeThumnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// makeThumnails2 고루틴을 이용하여 작업의 속도를 개선 했지만
// 이 함수에서 실행하는 고루틴이 완료되기 전에 main 루틴이 종료될 수 있음
func makeThumnails2(filenames []string) {
	for _, f := range filenames {
		go func(f string) {
			if _, err := thumbnail.ImageFile(f); err != nil {
				log.Println(err)
			}
		}(f)
	}
}

// makeThumnails3 채널을 이용하여 동기화, len(filenames) 만큼의 채널 이벤트를 확인한다
func makeThumnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// f 값을 명시적인 인자로 전달함.
		// 명시적으로 전달하지 않으면 외부 for 루프의 f를 캡쳐해서 사용하므로
		// 아래의 고루틴이 시작되는 시점에 f값이 부정확 하기 때문에, f 값을 명시적으로 전달
		go func(f string) {
			if _, err := thumbnail.ImageFile(f); err != nil {
				log.Println(err)
			}
			ch <- struct{}{}
		}(f)
	}
	// 고루틴들이 완료될 때까지 대기
	for range filenames {
		<-ch
	}
}

// makeThumnails4 지정된 파일들의 썸네일을 동시에 생성
// 중간에 실패하면 오류를 반환하도록
// 하지만 아래와 같은 방법으로 체크하면 err가 났을때 함수를 나가게 되고(errors 채널의 수신부 사라짐)
// 아직 처리하지 않은 나머지 고루틴에서, 해당 채널에 값을 보내려고 하고 있기 때문에,
// 고루틴 유출이 발생할 수 있다
func makeThumnails4(filenames []string) error {
	errors := make(chan error) // 버퍼 없는 채널
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: 오류: 고루틴 유출.
		}
	}
	return nil
}

// makeThumnails5 생성된 파일명을 임의의 순서로 반환. 중간에 실패하면 오류 반환
// 4버전관 달리 채널에 버퍼를 형성해, 고루틴이 메세지를 보낼 때 대기하지 않도록
func makeThumnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// makeThumnails6 채널에서 받은 각 파일의 썸네일을 생성
// 생성한 파일이 차지하는 바이트 수를 반환
func makeThumnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // 구동 중인 고루틴 수
	for f := range filenames {
		wg.Add(1)

		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer .. woker의 일이 모두 끝나면 sizes채널 close
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
