package main

/*
연습문제 8.12
새 클라이언트가 입장할 때마다 현재의 클라이언트 목록을 브로드캐스터가 알리게 하라. 이렇게 하려면 clients 목록이 필요하며,
entering과 leaving 채널에서 클라이언트의 이름도 기록해야 한다.

ch8
go run netcat3/netcat.go
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
)
type client struct {
	channel chan<- string // an outgoing message channel
	name    string
}

var (
	entering = make(chan *client) // 입장을 알리는 채널
	leaving  = make(chan *client) // 퇴장을 알리는 채널
	messages = make(chan string) // 메시지 채널
)

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
		// 새로운 커넥션을 기다리다가
		// 새로운 커넥션이 맺어지면 handleConn 고루틴 할당
		go handleConn(conn)
	}

}

// 현재의 클라이언트 목록을 브로드캐스터가 알리게 하라.
func broadcaster() {
	// 일종의 클라이언트 DB
	clients := make(map[string]*client)
	for {
		select {
		case msg := <-messages:
			// 메시지 브로드캐스팅
			for _, cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			// 새로운 접속자가 들어오면 기존 접속자의 리스트를 보내준다.
			go giveAllClients(cli.channel, clients)
			clients[cli.name] = cli

		case cli := <-leaving:
			// 기존 접속자가 나가는 거 방송
			delete(clients, cli.name)
			close(cli.channel) // 나간 사람의 채널 닫는다.
		}
	}
}


func giveAllClients(channel chan<- string, clients map[string]*client) {
	if len(clients) > 1 {
		channel <- "All clients:"
		for _, cli := range clients {
			channel <- cli.name
		}
	}
}

// 각 클라이언트마다 고루틴으로 돌아가는 함수
func handleConn(conn net.Conn) {
	ch := make(chan string)
	me := &client{channel: ch, name: conn.RemoteAddr().String()}
	go clientWriter(conn, ch)

	me.channel <- "You are " + me.name
	messages <- me.name + " has arrived"
	entering <- me

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- me.name + ": " + input.Text()
	}

	leaving <- me
	messages <- me.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

