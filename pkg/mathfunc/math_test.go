package mathfunc_test

import (
	"github.com/therenotomorrow/leetcode/pkg/mathfunc"
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{nums: []int{1, 3, 2, 4, 0}}, want: 4},
		{name: "negative", args: args{nums: []int{-1, -3, -2, -4, -5}}, want: -1},
		{name: "mixed", args: args{nums: []int{-1, 3, -2, 4, 0}}, want: 4},
		{name: "single", args: args{nums: []int{0}}, want: 0},
		{name: "zero", args: args{nums: []int{}}, want: math.MinInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Max(tt.args.nums...); got != tt.want {
				t.Errorf("Max() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{nums: []int{1, 3, 2, 4, 0}}, want: 0},
		{name: "negative", args: args{nums: []int{-1, -3, -2, -4, -5}}, want: -5},
		{name: "mixed", args: args{nums: []int{-1, 3, -2, 4, 0}}, want: -2},
		{name: "single", args: args{nums: []int{0}}, want: 0},
		{name: "zero", args: args{nums: []int{}}, want: math.MaxInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Min(tt.args.nums...); got != tt.want {
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
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{nums: []int{1, 3, 2, 4, 0}}, want: 10},
		{name: "negative", args: args{nums: []int{-1, -3, -2, -4, -5}}, want: -15},
		{name: "mixed", args: args{nums: []int{-1, 3, -2, 4, 0}}, want: 4},
		{name: "single", args: args{nums: []int{2}}, want: 2},
		{name: "zero", args: args{nums: []int{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathfunc.Sum(tt.args.nums...); got != tt.want {
				t.Errorf("Sum() = %v, want = %v", got, tt.want)
			}
		})
	}
}
