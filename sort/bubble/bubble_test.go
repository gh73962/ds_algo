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
			name: "test1",
			args: args{
				source: []int{5, 2, 1, 3, 4},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test2",
			args: args{
				source: []int{},
			},
			want: []int{},
		},
		{
			name: "test2",
			args: args{
				source: []int{1},
			},
			want: []int{1},
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
