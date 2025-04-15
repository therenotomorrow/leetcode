package golang

import "testing"

func TestMaxSubArrayLen(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{1, -1, 5, -2, 3}, k: 3}, want: 4},
		{name: "smoke 2", args: args{nums: []int{-2, -1, 2, 1}, k: 1}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxSubArrayLen(test.args.nums, test.args.k); got != test.want {
				t.Errorf("maxSubArrayLen() = %v, want = %v", got, test.want)
			}
		})
	}
}
