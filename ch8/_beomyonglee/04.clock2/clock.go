package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

/*
$ nc localhost 8000
21:28:05
21:28:06
21:28:07
21:28:08
*/
