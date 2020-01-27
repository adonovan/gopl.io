// 2와의 차이점은 뭔가?
// conn 이라는 스트림이 끝나더라도, 내부의 goroutine 이 끝날때까지
// main routine 이 기다려주게 하는 것

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- 0
	}()
	mustCopy(conn, os.Stdin)
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
