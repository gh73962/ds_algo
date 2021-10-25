package quick

import "log"

// 快速排序
// 选定pivot中心轴
// 将大于pivot的数字放在pivot的右边
// 小于pivot的在左边
// 将左右序列重复该操作
// 空间复杂度为o(1) 不额外申请空间

// 情况1 pivot是居中的
// 情况2 pivot在头
// 情况3 pivot在尾

func Quick(s []int) {
	l := len(s)
	if l <= 1 {
		return
	}
	mid := partition(s)
	Quick(s[:mid])
	Quick(s[mid+1:])
}

// 分区函数可以优化,优化pivot的选取方式,尽可能保证pivot在中间
func partition(s []int) (pos int) {
	log.Printf("origin is %+v", s)
	length := len(s)
	if length <= 1 {
		return
	}

	// 选定pivot中心轴
	var pivot, left, right int
	right = length - 1

	for {
		log.Printf("pivot is %d, left is %d, right is %d", s[pivot], s[left], s[right])
		var leftExchange, rightExchange bool

		// 寻找pivot的位置
		if left == right {
			s[left], s[pivot] = s[pivot], s[left]
			pos = left
			break
		}

		// 将大于pivot的数字放在pivot的右边
		if s[right] >= s[pivot] {
			right--
			continue
		} else {
			rightExchange = true
		}

		// 小于pivot的在左边
		if s[left] <= s[pivot] {
			left++
			continue
		} else {
			leftExchange = true
		}

		if leftExchange && rightExchange {
			s[left], s[right] = s[right], s[left]
		}

	}
	log.Printf("result is %+v", s)
	return
}
