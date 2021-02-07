package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

var count int = 1

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/getCount", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*func handle(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL)
}*/

func handle(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "%s = %s\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "request count: %d\n", count)
	mu.Unlock()
}
