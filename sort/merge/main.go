package main

import (
	"log"
)

func main() {
	arr := []int{5, 2, 3, 1, 6}
	arr = mergeSort(arr)
	log.Printf("result is %+v", arr)
}

//  归并排序
//  将排序元素最2分到最小单位，将最小单位进行排序，
//  然后进行merge，merge过程中再利用哨兵进行排序
// 其时间复杂度是非常稳定的，不管是最好情况、最坏情况，还是平均情况，时间复杂度都是 O(nlogn)
// 尽管每次合并操作都需要申请额外的内存空间，但在合并完成之后，临时开辟的内存空间就被释放掉了。
// 在任意时刻，CPU 只会有一个函数在执行，也就只会有一个临时的内存空间在使用。
// 临时内存空间最大也不会超过 n 个数据的大小，所以空间复杂度是 O(n)。

// 打散数组
func mergeSort(s []int) []int {
	log.Printf("s is %+v", s)
	length := len(s)
	// 终止条件
	if length <= 1 {
		return s
	}

	mid := length / 2
	log.Printf("mid is %+v", mid)
	arr1 := mergeSort(s[:mid])
	arr2 := mergeSort(s[mid:length])
	return merge(arr1, arr2)
}

// 合并数组
func merge(arr1, arr2 []int) []int {
	var result []int
	len1 := len(arr1)
	len2 := len(arr2)
	var i, j int

	for {
		if i < len1 && j < len2 {
			switch {
			// 保证稳定性
			case arr1[i] > arr2[j]:
				result = append(result, arr2[j])
				j++
			default:
				result = append(result, arr1[i])
				i++
			}
		} else if i < len1 {
			result = append(result, arr1[i])
			i++
		} else if j < len2 {
			result = append(result, arr2[j])
			j++
		} else {
			break
		}
	}
	return result
}
