package insertion

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
			name: "test1",
			args: args{
				origin: []int{1},
			},
			want: []int{1},
		},
		{
			name: "test2",
			args: args{
				origin: []int{5, 2, 3, 4, 1},
			},
			want: []int{1, 2, 3, 4, 5},
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
