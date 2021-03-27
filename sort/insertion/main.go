package main

import "log"

func main() {
	arr := []int{5, 1, 4, 2, 6, 3}
	log.Printf("origin arr is %+v", arr)
	arr = insertion(arr)
	log.Printf("arr is %+v", arr)
}

// 原地排序
// 稳定
// o(n2)
func insertion(s []int) []int {
	l := len(s)
	if l < 1 {
		return s
	}

	for i := 1; i < l; i++ {
		// i是标志位

		// 提取数据暂存
		v := s[i]
		// j是实际控制位置的
		j := i - 1

		for j >= 0 {
			// 寻找到要替换的位置
			if s[j] > v {
				s[j+1] = s[j]
			} else {
				break
			}
			//  此处j--调整位置
			j--
		}
		// 此处j+1不同于上面的j+1
		// 替换为暂存的数据
		s[j+1] = v
	}

	return s
}
