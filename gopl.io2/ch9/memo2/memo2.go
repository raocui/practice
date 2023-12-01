package memo2

// 扩展 Func类型和 (*Memo)Get 方法，让调用者可选择性地提供一个 done 通道方便取消操作(参考8.9节)。不要缓存被取消的 Func 调用结果。

type Func func(key string) (interface{}, error)


type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}



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
