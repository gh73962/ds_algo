package main

import "log"

// 选择排序
//  1.将当前值设为最小值
// 2.遍历未排序的元素,寻找真正的最小值,遍历后与最小值swap
// 不断交换元素位置,破坏了稳定性
// 时间复杂度 O(n2)
func main() {
	arr := []int{4, 1, 3, 5, 2}
	arr = selection(arr)
	log.Printf("result is %+v", arr)
}

func selection(s []int) []int {
	length := len(s)
	if length <= 1 {
		return s
	}

	for i := 0; i < length; i++ {
		min := s[i]
		var exchange int
		for j := i + 1; j < length; j++ {
			if min > s[j] {
				min = s[j]
				exchange = j
			}
		}
		if exchange > i {
			s[i], s[exchange] = s[exchange], s[i]
		}

	}

	return s
}
