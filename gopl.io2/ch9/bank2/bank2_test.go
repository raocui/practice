package bank2_test

import (
	"_/gopl.io2/ch9/bank2"
	"fmt"
	"testing"
)

func Test_Bank2(t *testing.T) {
	done := make(chan struct{})

	go func() {

		bank2.Deposit(100)
		done <- struct{}{}
	}()
	go func() {

		bank2.Deposit(100)
		done <- struct{}{}
	}()
	go func() {
		r := bank2.Withdraw(100)
		fmt.Printf("提现结果11：%v\n", r)
		done <- struct{}{}
	}()
	go func() {
		r := bank2.Withdraw(101)
		fmt.Printf("提现结果22：%v\n", r)
		done <- struct{}{}
	}()
	<-done
	<-done
	<-done
	m := bank2.Balance()
	fmt.Printf("余额：%d", m)
}
