package main

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	list, err := crawl("http://gopl.io")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", list)
}

// 入参： 一个url地址
// 出参： 输出该网页里的所有url地址
// 实现方案：获取该url的dom文档树， 广度优先遍历该树的每一个节点， 然后匹配是a连接的
func crawl(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("获取%s内容失败：%s", url, err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("不是返回200")
	}

	//将body转为文档树
	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析%s文档内容到文档树失败：%s", url, err.Error())
	}

}

func visit(links []string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key == "href" {
				links = append(links, v.Val)
			}
		}
	}

	for c := n.FirstChild; n != nil; c = n.NextSibling {
		
	}
}
