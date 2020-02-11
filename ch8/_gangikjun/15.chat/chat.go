package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string // 나가는 메시지 채널
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 수신된 모든 클라이언트 메시지
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // 모든 접속된 클라이언트
	for {
		select {
		case msg := <-messages:
			// 수신된 메시지를 모든 클라이언트로 나가는
			// 메시지 채널에 브로드 캐스트
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 나가는 클라이언트 메시지
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
