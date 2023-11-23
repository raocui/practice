package bank

// 数据竞态
// 对于一个如此简单的程序，我们一眼就可以看出，任意串行地调用 Deposit 和 Balanc
// 都可以得到正确的结果。即 Balance 会输出之前存人的金额总数。但如果这些函数的调用顺
// 序不是串行而是并行，Balance 就不保证输出正确结果了

var balance int

func Deposit(amount int) {

	balance = balance + amount
}
func Balance() int {
	return balance
}
