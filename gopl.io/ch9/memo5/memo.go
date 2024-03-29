// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

// !+Func
// Func是用于记忆的函数类型
// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// // result是调用 Func 的返回结果
// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently(随后) call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

// !+monitor goroutine 监控
// goroutine 监控 常常看、经常看！！！！！！！！
// map 变量限制在一个监控 goroutine 中，而Get的调用者则不得不改为发送消息。
// cache变量被限制在监控 goroutine (即(*Memo).server)中。监控 goroutine从request 通道中循环读取，直到该通道被 close 方法关闭。
// 对于每个请求，它先查询缓存如果没找到则创建并插人一个新的 entry。
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

//!-monitor
// 与基于互斥锁的版本类似，对于指定键的一次请求负责在该键上调用函数 f，保存结果到entry 中，最后通过关闭 ready 通道来广播准备完毕状态。
//这个流程通过(*entry).call实现
