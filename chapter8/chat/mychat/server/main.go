package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch chan string
	name string
}

var (
	entering = make(chan *client)
	leaving = make(chan *client)
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
	addr := conn.RemoteAddr().String()
	send := make(chan struct{})
	cli := &client{
		ch: make(chan string),
		name: addr,
	}
	go clientWrite(conn, cli)
	entering <- cli
	cli.ch <- "You are " + addr

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				conn.Close()
			case <-send:
				ticker.Reset(5 * time.Second)
			}
		}
	}()

	input := bufio.NewScanner(conn)

	for input.Scan() {
		send <- struct{}{}
		msg := addr + ": " + input.Text()
		message <- msg
	}

	defer func() {
		message <- cli.name + " is leaved."
		leaving <- cli
	}()
}

func clientWrite(conn net.Conn, cli *client) {
	for msg := range cli.ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcast() {
	cliMap := make(map[*client]bool)
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case cli := <-entering:
			cliMap[cli] = true
			for c := range cliMap {
				c.ch <- cli.name + " is arrived."
			}
		case cli := <-leaving:
			delete(cliMap, cli)
			close(cli.ch)
		case msg := <-message:
			for cli := range cliMap {
				select {
				case cli.ch <- msg:
					ticker.Reset(10 * time.Second)
				case <-ticker.C:
					break
				}
				cli.ch <- msg
			}
		}
	}
}
