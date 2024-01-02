package rotate

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		{args: args{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_rotate2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{args: args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{args: args{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}, want: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("rotate2() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
