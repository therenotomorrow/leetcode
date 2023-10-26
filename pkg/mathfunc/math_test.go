package mathfunc_test

import (
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		num  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{num: 1, nums: []int{1, 3, 2, 4, 0}}, want: 4},
		{name: "negative", args: args{num: -1, nums: []int{-1, -3, -2, -4, -5}}, want: -1},
		{name: "mixed", args: args{num: -1, nums: []int{-1, 3, -2, 4, 0}}, want: 4},
		{name: "single", args: args{num: 5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Max(tt.args.num, tt.args.nums...); got != tt.want {
				t.Errorf("Max() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		num  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{num: 1, nums: []int{1, 3, 2, 4, 0}}, want: 0},
		{name: "negative", args: args{num: -1, nums: []int{-1, -3, -2, -4, -5}}, want: -5},
		{name: "mixed", args: args{num: -1, nums: []int{-1, 3, -2, 4, 0}}, want: -2},
		{name: "single", args: args{num: 5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Min(tt.args.num, tt.args.nums...); got != tt.want {
				t.Errorf("Min() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{num: 5}, want: 5},
		{name: "negative", args: args{num: -5}, want: 5},
		{name: "zero", args: args{num: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Abs(tt.args.num); got != tt.want {
				t.Errorf("Abs() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		num  int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{num: 1, nums: []int{1, 3, 2, 4, 0}}, want: 11},
		{name: "negative", args: args{num: -1, nums: []int{-1, -3, -2, -4, -5}}, want: -16},
		{name: "mixed", args: args{num: -1, nums: []int{-1, 3, -2, 4, 0}}, want: 3},
		{name: "single", args: args{num: 5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Sum(tt.args.num, tt.args.nums...); got != tt.want {
				t.Errorf("Sum() = %v, want = %v", got, tt.want)
			}
		})
	}
}
