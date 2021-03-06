package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
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

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	defer conn.Close()
	for input.Scan() {
		go echo(conn, input.Text(), 1 * time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, shout)
	time.Sleep(delay)
	fmt.Fprintln(c, strings.ToLower(shout))
	time.Sleep(delay)
}
