package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var port = flag.String("port", "8010,8002,8003", "set the connected port which separated by comma")

func main() {
	flag.Parse()
	for {
		for _, p := range strings.Split(*port, ",") {
			connect(p)
		}
	}
}

func connect(port string) {
	conn, err := net.Dial("tcp", "localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(writer io.Writer, reader io.Reader) {
	_, err := io.Copy(writer, reader)
	if err != nil {
		log.Fatal(err)
	}
}
