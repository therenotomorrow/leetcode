package golang

import (
	"testing"
)

func TestMinimumSize(t *testing.T) {
	t.Parallel()

	type args struct {
		nums          []int
		maxOperations int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{9}, maxOperations: 2}, want: 3},
		{name: "smoke 2", args: args{nums: []int{2, 4, 8, 2}, maxOperations: 4}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumSize(test.args.nums, test.args.maxOperations); got != test.want {
				t.Errorf("minimumSize() = %v, want = %v", got, test.want)
			}
		})
	}
}
