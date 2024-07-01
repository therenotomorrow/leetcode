package golang

import "testing"

func TestMaxSubarrayLength(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 1, 2, 3, 1, 2}, k: 2}, want: 6},
		{name: "smoke 2", args: args{nums: []int{1, 2, 1, 2, 1, 2, 1, 2}, k: 1}, want: 2},
		{name: "smoke 3", args: args{nums: []int{5, 5, 5, 5, 5, 5, 5}, k: 4}, want: 4},
		{name: "test 730: wrong answer", args: args{nums: []int{3, 1, 1}, k: 1}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxSubarrayLength(test.args.nums, test.args.k); got != test.want {
				t.Errorf("maxSubarrayLength() = %v, want = %v", got, test.want)
			}
		})
	}
}
