/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package cache

import (
	"fmt"
	"os"
	"time"
)

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan result
	done chan struct{}
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-req.done:
			fmt.Printf("cancel request: %s\n", req.key)
			req.response <- result{value: make([]byte, 0), err: nil}
			break
		default:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key)
			}
			go e.delivery(req.response)
		}
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) delivery(resp chan<- result) {
	<- e.ready
	resp <- e.res
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	request := request{
		key:      key,
		response: make(chan result),
		done: done,
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	memo.requests<- request
	res := <-request.response
	return res.value, res.err
}
