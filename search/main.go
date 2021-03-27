package main

import "log"

func main() {
	s := []int{1, 2, 3, 4, 5, 5, 5, 10, 12, 15, 16, 23, 62, 122, 15123}
	pos := firstBigThan(s, 6)
	log.Printf("pos is %d\n\n", pos)
	// pos2 := recursion(s, 62)
	// log.Printf("pos2 is %d", pos2)
}

// 二分查找
// 二分查找针对的是一个有序的数据集合，查找思想有点类似分治思想。
// 每次都通过跟区间的中间元素对比，将待查找的区间缩小为之前的一半，直到找到要查找的元素，或者区间被缩小为 0。
// 时间复杂度o(logn)针对大数据极强
// 4 种常见二分查找的变形
// 1、查找第一个给定值的元素
// 2、查找最后一个给定值的元素
// 3、查找第一个>=给定值的元素
// 4、查找最后一个<=给定值的元素

// by 循环
// 寻找第一个K大的值
func firstEqual(s []int, value int) (pos int) {
	length := len(s)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)>>1

		switch {
		case value == s[mid]:
			if mid == 0 || s[mid-1] != value {
				return mid
			}
			high = mid - 1
		case value > s[mid]:
			low = mid + 1
		case value < s[mid]:
			high = mid - 1
		}
	}

	return -1
}

// 2、查找最后一个给定值的元素
func lastEqual(s []int, value int) int {
	length := len(s)
	low, high := 0, length-1
	for low <= high {
		mid := low + (high-low)>>1

		switch {
		case value == s[mid]:
			if mid == length-1 || s[mid+1] != value {
				return mid
			}
			low = mid + 1
		case value > s[mid]:
			low = mid + 1
		case value < s[mid]:
			high = mid - 1
		}
	}
	return -1
}

// 3、查找第一个>=给定值的元素
func firstBigThan(s []int, value int) int {
	length := len(s)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)>>1

		switch {
		case value <= s[mid]:
			if mid == 0 || s[mid-1] < value {
				return mid
			}
			high = mid - 1
		case value > s[mid]:
			low = mid + 1
		}
	}
	return -1
}

// 递归
func recursion(s []int, value int) (pos int) {
	return partition(s, value, 0, len(s)-1)
}

func partition(s []int, value, low, high int) int {

	if low > high {
		return -1
	}
	mid := low + (high-low)>>1

	switch {
	case value == s[mid]:
		return mid
	case value > s[mid]:
		return partition(s, value, mid+1, high)
	case value < s[mid]:
		return partition(s, value, low, mid-1)
	}

	return 0
}
