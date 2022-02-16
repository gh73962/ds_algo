package bubble

// 比较相邻的元素,将小的swap o(n^2) 稳定
func Bubble(source []int) []int {
	length := len(source)
	if length <= 1 {
		return source
	}

	for range source {
		var change bool
		for idx := 0; idx < length-1; idx++ {
			if change = source[idx] > source[idx+1]; change {
				source[idx], source[idx+1] = source[idx+1], source[idx]

			}
		}
		if !change {
			break
		}
	}
	return source
}
