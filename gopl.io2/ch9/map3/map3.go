package map3

import "image"

var icons map[string]image.Image

// 这个版本的 Icon 使用了延迟初始化:
func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("heartspng"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// 注意:并发不安全
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // 一次性地初始化
	}
	return icons[name]
}

func loadIcon(string) image.Image {
	return &image.RGBA{}
}

// 对于那些只被一个goroutine 访问的变量，上面的模式是没有问题的，但对于这个例子,并发调用 Icon 时这个模式就是不安全的。类似于银行例子中最早版本的 Deposit 函数,
// Icon 也包含多个步骤:检测 icons 是否为空，再加载图标，最后更新 icons 为一个非 nil值。
// 直觉可能会告诉你，竞态带来的最严重问题可能就是 loadIcons 函数会被调用多遍。当第一个goroutine 正忙于加载图标时，其他 goroutine 进人 Icon 函数，
// 会发现 icons 仍然是nil。所以仍然会调用 loadIcons。
// 但这个直觉仍然是错的(我希望你现在已经有一个关于并发的新直觉，那就是关于并发的直觉都不可靠)。回想一下9.4 节关于内存的讨论，在缺乏显式同步的情况下，编译器
// 和CPU在能保证每个goroutine 都满足串行一致性的基础上可以自由地重排访问内存的顺序。loadIcons一个可能的语句重排结果如下所示。它在填充数据之前把一个空 map 赋给
// icons:
// func loadIcons() {
// 	icons = make(map[string]image.Image)
// 	icons["spades.png"] = loadIcon("spades.png")
// 	icons["hearts.png"] = loadIcon("hearts.png")
// 	icons["diamonds.png"] = loadIcon("diamonds.png")
// 	icons["clubs.png"] = loadIcon("clubs.png")
// }

// 因此，一个goroutine 发现icons 不是nil并不意着变量的初始化肯定已经完成。
// 保证所有 goroutine 都能观察到 loadIcons 效果最简单的正确方法就是用一个互斥锁来做同步:

// var mu sync.Mutex // 保护 icons
// var icons map[string]image.Image
