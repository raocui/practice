// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// 避免数据竞态的方法是避免从多个 goroutine 访问同一个变量
// 把变量限制在单个goroutine内部。
// 由于其他 goroutine 无法直接访问相关变量，因此它们就必须使用通道来向受限
// goroutine 发送查询请求或者更新变量。这也是这句 Go 箴言的含义:“不要通过共享内存来通信，而应该通过通信来共享内存”。
//使用通道请求来代理一个受限变量的所有访问的 goroutine 称为该变量的监控 goroutine(monitor goroutine)。

// 即使一个变量无法在整个生命周期受限于单个 goroutine，加以限制仍然可以是解决并
// 发访问的好方法。比如一个常见的场景，可以通过借助通道来把共享变量的地址从上一步传
// 到下一步，从而在流水线上的多个 goroutine 之间共享该变量。在流水线中的每一步，在把
// 变量地址传给下一步后就不再访问该变量了，这样所有对这个变量的访问都是串行的。换个
// 说法，这个变量先受限于流水线的一步，再受限于下一步，以此类推。这种受限有时也称为
// 串行受限。
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			fmt.Println(balance)
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
