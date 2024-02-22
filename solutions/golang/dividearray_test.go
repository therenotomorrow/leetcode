package golang

import (
	"reflect"
	"testing"
)

func TestDivideArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1}, k: 2},
			want: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{name: "smoke 2", args: args{nums: []int{1, 3, 3, 2, 7, 3}, k: 3}, want: [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideArray() = %v, want = %v", got, tt.want)
			}
		})
	}
}
