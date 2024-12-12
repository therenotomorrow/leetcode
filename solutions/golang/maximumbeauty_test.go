package golang

import (
	"testing"
)

func TestMaximumBeauty(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{4, 6, 1, 2}, k: 2}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1, 1}, k: 10}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumBeauty(test.args.nums, test.args.k); got != test.want {
				t.Errorf("maximumBeauty() = %v, want = %v", got, test.want)
			}
		})
	}
}
