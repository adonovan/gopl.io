package main

// 서버가 연결마다 sync.WaitGroup을 사용해 활성화된 echo 고루틴 개수를 세도록 수정
// 활성화 수가 0이 되면 TCP 쓰기를 닫아라

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}

	// 활성화된 echo가 없다면 tcp write close
	go func() {
		wg.Wait()
		tc, ok := c.(*net.TCPConn)
		if ok {
			tc.CloseWrite()
		}
	}()

	c.Close()
}

func main() {
	lintener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lintener.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}
