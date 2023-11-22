package solutions

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, want: []int{1, 4, 2, 7, 5, 3, 8, 6, 9}},
		{name: "smoke 2", args: args{nums: [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}}, want: []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want = %v", got, tt.want)
			}
		})
	}
}
