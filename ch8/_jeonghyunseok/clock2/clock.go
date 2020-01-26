// 클락1이랑 뭐가 다를까? 그냥 defer 유무인듯 하다.package clock2
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer func() {
		c.Close()
		log.Print("disconneted")
	}()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Print("connected")
		go handleConn(conn)
	}
}

/*
go run .
go run ..\netcat1.go

*/
