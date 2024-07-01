package golang

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{10, 5, 2, 6}, k: 100}, want: 8},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3}, k: 0}, want: 0},
		{name: "test 97: wrong answer", args: args{nums: []int{1, 2, 3, 4, 5}, k: 1}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numSubarrayProductLessThanK(test.args.nums, test.args.k); got != test.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want = %v", got, test.want)
			}
		})
	}
}
