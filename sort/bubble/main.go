package main

import "fmt"

func main() {
	arr := []int{8, 5, 2, 3, 1, 6, 4, 9, 7, 5}
	arr = bubble(arr)
	fmt.Printf("arr is %+v", arr)
}

// 冒泡排序
// 比较相邻的元素,将小的swap
// 优化点 当无数据交换时,则不用再进行排序
// o(n2)
// 稳定,不破坏原来的顺序
func bubble(source []int) []int {
	length := len(source)
	if length <= 1 {
		return source
	}
	for i := 0; i < length; i++ {
		var flag bool
		for j := 0; j < length-1; j++ {
			if source[j] > source[j+1] {
				source[j], source[j+1] = source[j+1], source[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return source
}
