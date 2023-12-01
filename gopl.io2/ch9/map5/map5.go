package map5

import (
	"image"
	"sync"
)

func loadIcon(string) image.Image {
	return &image.RGBA{}
}
func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("heartspng"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// 采用互斥锁访问 icons 的额外代价是两个 goroutine 不能并发访问这个变量，即使在变量已经安全完成初始化且不再更改的情况下，也会造成这个后果。
//使用一个可以并发读的锁就可以改善这个问题:

var mu sync.RWMutex // 保护 icons
var icons map[string]image.Image

// 并发安全
func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()
	// 获取互斥锁
	mu.Lock()
	if icons != nil { // 注意:必须重新检查 nil值
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}

// 这里有两个临界区域。goroutine 首先获取一个读锁，查阅 map，然后释放这个读锁。如果条目能找到(常见情况)，就返回它。
// 如果条目没找到，goroutine 再获取一个写锁。由于不先释放一个共享锁就无法直接把它升级到互斥锁，
// 为了避免在过渡期其他 goroutine 已经初始化了 icons，所以我们必须重新检查 nil 值。

// 上面的模式具有更好的并发性，但它更复杂并且更容易出错。幸运的是，sync 包提供了
// 针对一次性初始化问题的特化解决方案:sync.once。
