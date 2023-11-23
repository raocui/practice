// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	bank "_/gopl.io2/ch9/bank1"
	"fmt"
	"testing"
)

// 注意：Bob 的存款在 Alice 的存款操作中间执行
// 晚于账户余额读取 (balance+amount)但早于余额更新(balance =...)，这会导致 Bob 存的
// 钱消失了。这是因为 Alice 的存款操作 A1 实际上是串行的两个操作，读部分和写部分，我们
// 称之为 A1r 和 A1w。
func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)                //A1
		fmt.Println("=", bank.Balance()) //A2
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100) //B
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
