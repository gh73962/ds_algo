package leetcode

import (
	"log"
	"testing"
)

func TestFindBySum(t *testing.T) {
	type args struct {
		array []int
		sum   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				array: []int{600, 4200, 1800, 8200, 18900, 11900, 19300, 20500, 67600},
				sum:   78300,
			},
		},
		{
			name: "2",
			args: args{
				array: []int{600, 3600, 4200, 1800, 8200, 18900, 11900, 19300, 28100, 32900, 20500, 67600},
				sum:   70900,
			},
		},
		// {
		// 	name: "2",
		// 	args: args{
		// 		array: []int{1, 2, 3, 4, 5},
		// 		sum:   5,
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindBySum(tt.args.array, tt.args.sum)
		})
	}
}

func TestPermutation(t *testing.T) {
	type args struct {
		array []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				array: []int{1, 2, 3, 4, 5},
				n:     3,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := Permutation(tt.args.array, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Permutation() = %v, want %v", got, tt.want)
			// }
			log.Printf("result=%+v", Combinations(tt.args.array, tt.args.n))
		})
	}
}
