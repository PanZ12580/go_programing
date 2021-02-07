package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	addr := conn.RemoteAddr()
	ch := make(chan string)
	go clientWrite(conn, ch)
	entering <- ch
	ch <- "You are " + addr.String()
	message <- addr.String() + " is arrived."

	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg := addr.String() + ": " + input.Text()
		message <- msg
	}

	ch <- addr.String() + " is leaved."
	leaving <- ch
}

func clientWrite(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcast() {
	cliMap := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			cliMap[cli] = true
		case cli := <-leaving:
			delete(cliMap, cli)
			close(cli)
		case msg := <-message:
			for cli := range cliMap {
				cli <- msg
			}
		}
	}
}
