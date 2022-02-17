package selection

// 选择排序 (循环找最小替换)
// 1.将当前值设为最小值
// 2.遍历未排序的元素,寻找真正的最小值,遍历后与最小值swap
// 不断交换元素位置,破坏了稳定性
// 时间复杂度 O(n2)

func Sort(origin []int) []int {
	l := len(origin)
	if l <= 1 {
		return origin
	}
	for i := 0; i < l; i++ {
		var (
			minIndex = i
			exchange bool
		)
		for start := i; start < l; start++ {
			if origin[start] < origin[minIndex] {
				minIndex = start
				exchange = true
			}
		}
		if exchange {
			origin[i], origin[minIndex] = origin[minIndex], origin[i]
		}
	}
	return origin
}
