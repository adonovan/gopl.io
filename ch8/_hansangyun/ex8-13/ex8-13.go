package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

/*
연습문제 8.13
채팅 서버가 마지막 5분 안에 어떤 메시지도 전송하지 않는 등의 유휴 클라이언트의 접속을 해제하게 하라.

힌트: 다른 고루틴에서 conn.Close()를 호출하면 input.Scan() 등에 의해 활성화된 Read호출로 인한 대기 상태를 해제한다.
*/

// 타임아웃 시간
const timeout = 10 * time.Second

type client struct {
	Out  chan<- string // an outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Out <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.Out <- "Present:"
			for c := range clients {
				cli.Out <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

// 각 클라이언트마다 고루틴으로 돌아가는 함수
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	/*
	바뀐 점 :
	타이머를 두고 메시지를 입력할 때마다 타임아웃 리셋
	 */
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()

		// 타이머 리셋
		timer.Reset(timeout)
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 개별 커넥션(클라이언트)와 상관없이 브로드캐스팅을 위한 고루틴
	go broadcaster()


	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// 새로운 커넥션을 대기하다가
		// 새로운 커넥션이 맺어지면 handleConn 고루틴 할당
		go handleConn(conn)
	}
}
