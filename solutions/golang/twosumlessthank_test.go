package golang

import "testing"

func TestTwoSumLessThanK(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{34, 23, 1, 24, 75, 33, 54, 8}, k: 60}, want: 58},
		{name: "smoke 2", args: args{nums: []int{10, 20, 30}, k: 15}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := twoSumLessThanK(test.args.nums, test.args.k); got != test.want {
				t.Errorf("twoSumLessThanK() = %v, want = %v", got, test.want)
			}
		})
	}
}
