package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8080", "set the listened port")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05") + "\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
