package leetcode

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		list1: &ListNode{
		// 			Val:  1,
		// 			Next: nil,
		// 		},
		// 		list2: &ListNode{
		// 			Val:  1,
		// 			Next: nil,
		// 		},
		// 	},
		// 	want: nil,
		// },
		{
			name: "test2",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListPrint(mergeTwoLists(tt.args.list1, tt.args.list2))
		})
	}
}
