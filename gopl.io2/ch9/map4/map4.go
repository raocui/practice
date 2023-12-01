package map4

import (
	"image"
	"sync"
)

var mu sync.Mutex // 保护 icons
var icons map[string]image.Image

// 并发安全
func Icon(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

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

// 采用互斥锁访问 icons 的额外代价是两个 goroutine 不能并发访问这个变量，即使在变量已经安全完成初始化且不再更改的情况下，也会造成这个后果。
//使用一个可以并发读的锁就可以改善这个问题:
