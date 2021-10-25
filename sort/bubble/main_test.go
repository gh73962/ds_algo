package bubble

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{[]int{8, 5, 2, 3, 1, 6, 4, 9, 7, 5}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bubble(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
