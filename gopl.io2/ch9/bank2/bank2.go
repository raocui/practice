package bank2

//
// 向 gopl.io/ch9/bank1 程序添加一个函数 withdraw(amount int)bool。结果应当反
// 映交易成功还是由于余额不足而失败。函数发送到监控 goroutine 的消息应当包含取款金额和
// 一个新的通道，这个通道用于监控 goroutine 把布尔型的结果发送回 withdraw 函数。
// 数据竞态： 
//使用通道请求来代理一个受限变量的所有访问的 goroutine 称为该变量的监控 goroutine(monitor goroutine)

var deposits = make(chan int)  // send amount to deposit
var balances = make(chan int)  // receive balance
var withdraws = make(chan int) // 发送amount 到 withdraws（发送提现金额到提现通道）
var result = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			//fmt.Println(balance)
		case amount := <-withdraws:
			balance -= amount
			if balance >= 0 {
				result <- true
			} else {
				balance += amount
				result <- false

			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
