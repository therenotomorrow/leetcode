package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := golang.Max(tt.args.nums...); got != tt.want {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := golang.Min(tt.args.nums...); got != tt.want {
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
			if got := golang.Abs(tt.args.num); got != tt.want {
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
			if got := golang.Sum(tt.args.nums...); got != tt.want {
				t.Errorf("Sum() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestManhattan(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{x1: 1, y1: 2, x2: 3, y2: 4}, want: 4},
		{name: "negative", args: args{x1: -2, y1: -3, x2: -5, y2: -7}, want: 7},
		{name: "mixed", args: args{x1: -1, y1: 2, x2: 3, y2: -4}, want: 10},
		{name: "zero", args: args{x1: -1, y1: 2, x2: 3, y2: -4}, want: 10},
		{name: "horizontal", args: args{x1: 0, y1: 0, x2: 5, y2: 0}, want: 5},
		{name: "vertical", args: args{x1: 0, y1: 0, x2: 0, y2: 8}, want: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := golang.Manhattan(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("Manhattan() = %v, want = %v", got, tt.want)
			}
		})
	}
}
