package main

import (
	"fmt"
)

// 括号匹配的检验： 检验括号是否匹配的方法用“期待的急迫程度”这个概念来描述
// 分析可能出现的不匹配的情况:
// 到来的右括弧非是所“期待”的
// 到来的是“不速之客”
// 直到结束，也没有到来所“期待的
// 正确的格式： ([]()) 或 [([][])]
func main() {

}

func matching(str string) error {
	sli := []rune(str)
	data := make([]string, 0)
flag:
	for _, v := range sli {
		c := string(v)
		switch c {
		case "(", "[", "{":
			data = append(data, c)
		case ")", "]", "}":
			if len(data) < 1 {
				return fmt.Errorf("不速之客， 右边的括号多于左边")
			}
			pop := data[len(data)-1]
			data = data[:len(data)-1]
			switch c {
			case ")":
				if pop != "(" {
					goto flag
				}
			case "]":
				if pop != "[" {
					goto flag
				}
			case "}":
				if pop != "{}" {
					goto flag
				}
			}

		default:
			return fmt.Errorf("非法括号：%d", c)

		}
	}

	if len(data)>0{}

}
