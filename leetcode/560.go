package leetcode

import (
	"log"
	"math/bits"
)

func FindBySum(array []int, sum int) {
	length := len(array)

	for i := 2; i < length; i++ {
		subsets := Combinations(array, i)
		for _, sets := range subsets {
			var temp int
			for _, v := range sets {
				temp += v
			}
			if temp == sum {
				log.Printf("answer:%+v", sets)
			}
		}
	}
}

func Combinations(set []int, n int) (subsets [][]int) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
