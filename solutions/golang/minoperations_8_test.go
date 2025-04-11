package golang

import "testing"

func TestMinOperations8(t *testing.T) {
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
		{name: "smoke 1", args: args{nums: []int{5, 2, 5, 4, 5}, k: 2}, want: 2},
		{name: "smoke 2", args: args{nums: []int{2, 1, 2}, k: 2}, want: -1},
		{name: "smoke 3", args: args{nums: []int{9, 7, 5, 3}, k: 1}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations8(test.args.nums, test.args.k); got != test.want {
				t.Errorf("minOperations8() = %v, want = %v", got, test.want)
			}
		})
	}
}
