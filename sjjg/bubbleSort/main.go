package main

import "fmt"

func main() {
	data := []int{4, 5, 4, 6, 4, 7, 2, 3, 1}
	fmt.Println(BubbleSort(data))
}

func BubbleSort(data []int) []int {
	length := len(data)

	for j := length - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if data[i] > data[i+1] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
