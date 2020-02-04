package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8080, "listen port")

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
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
nc localhost 8080
21:33:12
21:33:13
21:33:14
21:33:15
21:33:16
*/
