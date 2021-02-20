/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package cache

import "sync"

type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{}
}

type Memo struct {
	f Func
	cache map[string]*entry
	mu sync.Mutex
}

type Func func(string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{f : f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		e.res.value, e.res.err = memo.f(key)
		memo.cache[key] = e
		memo.mu.Unlock()
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
