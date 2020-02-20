package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

/*
연습문제 8.14
각 클라이언트가 입장 시 이름을 제공하게 채팅 서버의 네트워크 프로토콜을 변경하라.
각 메시지에 붙이는 송신자의 식별자로 네트워크 주소 대신 이 이름을 사용하라.
*/


const timeout = 30 * time.Second

type client struct {
	Out  chan<- string // outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.Out <- msg:
				default:
					// Skip client if it's reading messages slowly.
				}
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

func handleConn(conn net.Conn) {
	out := make(chan string, 10) // outgoing client messages
	go clientWriter(conn, out)
	in := make(chan string) // incoming client messages
	go clientReader(conn, in)

	var who string
	nameTimer := time.NewTimer(timeout)
	out <- "Enter your name:"
	select {
	case name := <-in:
		who = name
	case <-nameTimer.C:
		conn.Close()
		return
	}
	cli := client{out, who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	idle := time.NewTimer(timeout)

Loop:
	for {
		select {
		case msg := <-in:
			messages <- who + ": " + msg
			idle.Reset(timeout)
		case <-idle.C:
			conn.Close()
			break Loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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
