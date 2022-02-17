package insertion

// Sort 原地排序 稳定 o(n2)
func Sort(origin []int) []int {
	l := len(origin)
	if l <= 1 {
		return origin
	}
	for unsortedHead := 1; unsortedHead < l; unsortedHead++ {
		var (
			unsortedValue = origin[unsortedHead]
			insertIndex   = unsortedHead - 1
		)
		for ; insertIndex >= 0; insertIndex-- {
			if origin[insertIndex] > unsortedValue {
				origin[insertIndex+1] = origin[insertIndex]
			} else {
				break
			}
		}
		origin[insertIndex+1] = unsortedValue
	}
	return origin
}
