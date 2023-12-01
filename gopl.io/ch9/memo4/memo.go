// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 276.

// Package memo provides a concurrency-safe memoization a function of
// a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// !+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// 在下面的 Memo 版本中，map 的每个元素是一个指向entry 结构的指针。除了
// 与之前一样包含一个已经记住的函数f 调用结果之外，每个entry 还新加了一个通道 ready
// 在设置entry的result字段后，通道会关闭，正在等待的 goroutine 会收到广播(参考8.9
// 节)，然后就可以从 entry 读取结果了。

// 使用一个斥量来保护被多个调用Get 的 goroutine 访问的 map 变量
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

// 现在调用Get 会先获取保护 cache map 的互斥锁，再从map 中查询一个指向已有entry 的指针，如果没有查找到，就分配并插人一个新的 entry，最后释放锁。
// 如果要查询的entry存在，那么它的值可能还没准备好(另外一个goroutine 有可能还在调用慢函数 f)所以主调 goroutine就需要等待 entry 准备好才能读取entry 中的 result 数据，
// 具体的实现方法就是从 ready 通道读取数据，这个操作会一直阻塞到通道关闭。
// 如果要查询的entry 不存在，那么当前的 goroutine 就需要新插人一个没有准备好的entry到map 里，并负责调用慢函数 f，更新entry，最后向其他正在等待的 goroutine 广播 数据已准备完毕的消息。
// 注意，entry中的变量e.res.value 和e.res.err 被多个goroutine共享。创建entry的goroutine 设置了这两个变量的值，其他 goroutine 在收到数据准备完毕的广播后开始读这两个变量。
// 尽管变量被多个 goroutine 访问，但此处不需要加上互斥锁。ready 通道的关闭
// 先于其他goroutine收到广播事件，所以第一个goroutine 的变量写入事件也先于后续多个goroutine的读取事件。在这个情况下数据竞态不存在。
// 这里的并发、重复抑制、非阻塞缓存就完成了。
//!-
