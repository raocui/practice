package map6

import (
	"image"
	"sync"
)

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("heartspng"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}
func loadIcon(string) image.Image {
	return &image.RGBA{}
}

// 幸运的是，sync 包提供了针对一次性初始化问题的特化解决方案:sync.once。
// 从概念上来讲，once 包含一个布尔变量和一个互斥量，布尔变量记录初始化是否已经完成，互斥量则负责保护这个布尔变量和客户端的数据结构。
// once 的唯一方法 Do 以初始化函数作为它的参数。让我们看一下 once 简化后的 Icon 函数:
var loadIconsOnce sync.Once
var icons map[string]image.Image

// 并发安全
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

//每次调用 Do(loadIcons)时会先锁定互斥量并检查里边的布尔变量。在第一次调用时,这个布尔变量为假，
// Do 会调用 loadicons 然后把变量设置为真。后续的调用相当于空操作,只是通过互斥量的同步来保证 loadIcons 对内存产生的效果 (在这里就是 icons 变量)对所
// 有的 goroutine可见。以这种方式来使用 sync.once，可以避免变量在正确构造之前就被其他goroutine 分享。
