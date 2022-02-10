package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	var (
		node1 = ListNode{Val: 1}
		node2 = ListNode{Val: 2}
		// node3 = ListNode{Val: 3}
		// node4 = ListNode{Val: 4}
		// node5 = ListNode{Val: 5}
	)
	node1.Next = &node2
	node2.Next = &node1
	// node3.Next = &node4
	// node4.Next = &node5
	// node5.Next = &node2

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				head: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
