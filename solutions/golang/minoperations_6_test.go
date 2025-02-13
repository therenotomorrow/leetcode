package golang

import "testing"

func TestMinOperations6(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{2, 11, 10, 1, 3}, k: 10}, want: 2},
		{name: "smoke 2", args: args{nums: []int{1, 1, 2, 4, 9}, k: 20}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations6(test.args.nums, test.args.k); got != test.want {
				t.Errorf("minOperations6() = %v, want = %v", got, test.want)
			}
		})
	}
}
