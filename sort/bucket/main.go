package main

import "log"

func main() {
	arr := []int{}

	log.Printf("final is %+v", arr)
}

// 桶排序
// 将数据划分到多个有序的桶里,再按序将桶里数据提取出来
// 难点:如何划分有序桶,等分效果不好
// 适合于外部排序
// 时间复杂度O(n)
func bucket(s []int) {

}
