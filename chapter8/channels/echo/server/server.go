package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
	texts := make(chan string)
	for {
		go func() {
			if input.Scan() {
				texts <- input.Text()
			} else {
				fmt.Println("over")
				return
			}
		}()

		select {
		case <-time.After(10 * time.Second):
			fmt.Println("close connection")
			return
		case t := <-texts:
			go echo(conn, t, 1 * time.Second)
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	wg.Add(1)
	defer wg.Done()

	go func() {
		wg.Wait()
		if tc, ok := c.(*net.TCPConn); ok {
			tc.CloseWrite()
		}
	}()

	fmt.Fprintln(c, strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, shout)
	time.Sleep(delay)
	fmt.Fprintln(c, strings.ToLower(shout))
	time.Sleep(delay)
}
