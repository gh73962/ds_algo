package selection

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		origin []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				origin: []int{1},
			},
			want: []int{1},
		},
		{
			name: "2",
			args: args{
				origin: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "3",
			args: args{
				origin: []int{4, 1, 3, 2},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
