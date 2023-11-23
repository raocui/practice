package map1

import (
	"fmt"
)

// 数据竞态(data race)发生于两个goroutine 并发读写同一个变量并且至少其中一个是写入时。（数据竞态的定义）
// 变量icons 被多个goroutine 并发读写，发生了数据竞态
var icons = make(map[string]string)

func loadIcon(name, flag string) string {
	return fmt.Sprintf("%s_%s", name, flag)
}

// 注意:并发不安全
func Icon(name, flag string) string {
	fmt.Printf("map变量icons的地址是:%p \n", icons)
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name, flag)
		icons[name] = icon
	}
	return icon
}
func GetIcons() {
	fmt.Printf("map变量的结果:%#v", icons)
}
