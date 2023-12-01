package map2

import "image"

// 如果在创建其他 goroutine 之前就用完整的数据来初始化 map，并且不再修改。
// 那么无论多少goroutine 也可以安全地并发调用 Icon，因为每个goroutine 都只读取这个map.
// 预先初始化一个变量
var icons = map[string]image.Image{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// 并发安全, 该icons变量， 在main函数执行之前就已经完成了初始化， 之后不在改动， 只读是并发安全的。
func Icon(name string) image.Image { return icons[name] }

func loadIcon(name string) image.Image {
	return nil
}
