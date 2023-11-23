// Package bank provides a concurrency-safe single-account bank.
package bank3

//!+
import "sync"

// Bob 的 100 美元存款消失了，没留下任何线索，Bob 感到很焦虑，为了解决这个问题
// Bob 写了一个程序，每秒钟查询数百次他的账户余额。这个程序同时在他家里、公司里和
// 他的手机上运行。银行注意到快速增长的业务请求正在拖慢存款和取款操作，因为所有的
// Balance 请求都是串行运行的，持有互斥锁并暂时妨碍了其他goroutine 运行。
// 因为 Balance 函数只须读取变量的状态，所以多个 Balance 请求其实可以安全地并发运
// 行，只要 Deposit 和 withdraw 请求没有同时运行即可。在这种场景下，我们需要一种特殊类
// 型的锁，它允许只读操作可以并发执行，但写操作需要获得完全独享的访问权限。这种锁称
// 为多读单写锁，Go语言中的 sync.RWMutex 可以提供这种功能:
var (
	mu      sync.RWMutex // // 读写互斥锁:sync.RWMutex guards balance
	balance int
)

// Balance 函数现在可以调用 RLock 和 RUnlock 方法来分别获取和释放一个读锁 (也称为共享锁)。
//Deposit 函数无须更改，它通过调用 mu.Lock 和 mu.Unlock 来分别获取和释放一个写锁(也称为互斥锁)。
// 经过上面的修改之后，Bob 的绝大部分 Balance 请求可以并行运行且能更快完成。因此
// 锁可用的时间比例会更大，Deposit 请求也能得到更及时的响应。
func Deposit(amount int) {
	mu.Lock() //写锁(也称为互斥锁)
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.RLocker() //读锁(共享锁)
	b := balance
	mu.RUnlock()
	return b
}

//!-
