package main

import "fmt"

func main() {
	data := []int{4, 5, 4, 6, 4, 7, 2, 3, 1}
	fmt.Println(quickSort(data, 0, 8))
}

// 快速排序
func quickSort(data []int, l, r int) []int {

	if l < r {
		i := l
		j := r
		x := data[i]

		for i < j {
			for i < j && data[j] > x {
				j--
			}
			if i < j {
				data[i] = data[j]
				i++
			}
			for i < j && data[i] < x {
				i++
			}
			if i < j {
				data[j] = data[i]
				j--
			}
			data[i] = x
			quickSort(data, l, i-1)
			quickSort(data, i+1, r)
		}

	}
	return data
}
