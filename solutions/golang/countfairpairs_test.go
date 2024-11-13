package golang

import (
	"testing"
)

func TestCountFairPairs(t *testing.T) {
	t.Parallel()

	type args struct {
		nums  []int
		lower int
		upper int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 7, 4, 4, 5}, lower: 3, upper: 6}, want: 6},
		{name: "smoke 2", args: args{nums: []int{1, 7, 9, 2, 5}, lower: 11, upper: 11}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countFairPairs(test.args.nums, test.args.lower, test.args.upper); got != test.want {
				t.Errorf("countFairPairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
