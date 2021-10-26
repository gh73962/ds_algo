package radix

import (
	"log"
	"strconv"
)

// 基数排序
// 例如排数字，则旭需有0-9 10个桶，（MSD最高位优先法和最低位优先LSD）按位分配到桶，再按桶序取出来，再进行下一位
// 若是字符串,则是26个桶
// LSD的基数排序适用于位数小的数列，如果位数多的话，使用MSD的效率会比较好。
// n个记录，d个关键码，关键码的取值范围为radix, O(d(n+radix)), radix是常数则为o(n)
func Int(s []int) []int {
	return msd(s, maxDigit(s))
}

func maxDigit(s []int) int {
	var max int

	for _, value := range s {
		if value > max {
			max = value
		}
	}
	return len(strconv.Itoa(max))
}

func msd(s []int, max int) []int {
	var result []int
	for max > 0 {
		result = []int{}
		var buckets = make([][]int, 10)
		for _, value := range s {

			index := value / getBase(max) % 10
			buckets[index] = append(buckets[index], value)
		}
		for _, bucket := range buckets {
			for _, value := range bucket {
				result = append(result, value)
			}
		}
		log.Printf("result is %+v", result)
		s = result
		max--
	}

	return result
}

func getBase(maxDigit int) int {
	if maxDigit == 0 {
		return 0
	}
	result := 1
	maxDigit--
	for maxDigit > 0 {
		result *= 10
		maxDigit--
	}
	return result
}
