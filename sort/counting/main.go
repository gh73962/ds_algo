package main

import "log"

func main() {
	arr := []int{2, 5, 3, 0, 2, 3, 0, 3}

	log.Printf("final is %+v", counting(arr))
}

// 计数排序
// 将原始数据A划分到统计数组B,B的下标是一个统计指标,值是个数;将B的元素各个相加结果存在C
// 结合AB处理到结果集B(从后往前遍历原始数组,从C中匹配对应下标的值,获得的值就是排序后的index)
func counting(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	countingArr := make([]int, getMaxValue(s))
	// 统计
	for _, value := range s {
		countingArr[value]++
	}
	log.Printf("after count,arr is %+v", countingArr)
	// 求和
	for index := range countingArr {
		if index == 0 {
			continue
		}
		countingArr[index] = countingArr[index] + countingArr[index-1]
	}
	log.Printf("after sum,arr is %+v", countingArr)
	// 倒序查找
	result := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		value := s[i]
		pos := countingArr[value] - 1 // 注意越界
		// 获取位置后，原位置的统计数--
		countingArr[value]--

		result[pos] = value
	}
	return result
}

func getMaxValue(s []int) (max int) {
	for index, value := range s {
		if index == 0 {
			max = value
		}
		if value > max {
			max = value
		}
	}
	max++
	log.Printf("max is %+v", max)
	return
}
