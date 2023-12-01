// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 276.

// Package memo provides a concurrency-safe memoization a function of
// type Func.  Requests for different keys run concurrently.
// Concurrent requests for the same key result in duplicate work.
package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// !+
// 性能再度得到提升，但我们注意到某些 URL 被获取了两次。在两个或者多个 goroutine
// 几乎同时调用Get 来获取同一个URL时就会出现这个问题。两个 goroutine 都首先查询缓
// 存，发现缓存中没有需要的数据，然后调用那个慢函数 f，最后又都用获得的结果来更新
// map，其中一个结果会被另外一个覆盖。
// 在理想情况下我们应该避免这种额外的处理。这个功能有时称为重复抑制 (duplicate suppression)。
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

//!-
