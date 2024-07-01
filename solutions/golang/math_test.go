package golang_test

import (
	"testing"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func TestMax(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := golang.Max(test.args.nums...); got != test.want {
				t.Errorf("Max() = %v, want = %v", got, test.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := golang.Min(test.args.nums...); got != test.want {
				t.Errorf("Min() = %v, want = %v", got, test.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := golang.Abs(test.args.num); got != test.want {
				t.Errorf("Abs() = %v, want = %v", got, test.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := golang.Sum(test.args.nums...); got != test.want {
				t.Errorf("Sum() = %v, want = %v", got, test.want)
			}
		})
	}
}

func TestManhattan(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := golang.Manhattan(test.args.x1, test.args.y1, test.args.x2, test.args.y2); got != test.want {
				t.Errorf("Manhattan() = %v, want = %v", got, test.want)
			}
		})
	}
}
