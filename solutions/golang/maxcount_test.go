package golang

import (
	"testing"
)

func TestMaxCount(t *testing.T) {
	t.Parallel()

	type args struct {
		banned []int
		n      int
		maxSum int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{banned: []int{1, 6, 5}, n: 5, maxSum: 6}, want: 2},
		{name: "smoke 2", args: args{banned: []int{1, 2, 3, 4, 5, 6, 7}, n: 8, maxSum: 1}, want: 0},
		{name: "smoke 3", args: args{banned: []int{11}, n: 7, maxSum: 50}, want: 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxCount(test.args.banned, test.args.n, test.args.maxSum); got != test.want {
				t.Errorf("maxCount() = %v, want = %v", got, test.want)
			}
		})
	}
}
