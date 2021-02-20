package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	done := make(chan struct{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	if tc, ok := conn.(*net.TCPConn); ok {
		tc.CloseWrite()
	}
	<- done
}

func mustCopy(writer io.Writer, reader io.Reader) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		log.Fatal(err)
	}
}
