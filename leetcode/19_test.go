package leetcode

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: NewHead(5),
				n:    1,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				head: NewHead(5),
				n:    2,
			},
			want: nil,
		},
		{
			name: "test3",
			args: args{
				head: NewHead(5),
				n:    3,
			},
			want: nil,
		},
		{
			name: "test4",
			args: args{
				head: NewHead(5),
				n:    4,
			},
			want: nil,
		},
		{
			name: "test50",
			args: args{
				head: NewHead(50),
				n:    5,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListPrint(RemoveNthFromEnd(tt.args.head, tt.args.n))
		})
	}
}

func TestNewHead(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				length: 1,
			},
			want: nil,
		},
		{
			name: "5",
			args: args{
				length: 5,
			},
			want: nil,
		},
		{
			name: "10",
			args: args{
				length: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListPrint(NewHead(tt.args.length))
		})
	}
}
