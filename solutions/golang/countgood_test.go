package golang

import "testing"

func TestCountGood(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{nums: []int{1, 1, 1, 1, 1}, k: 10}, want: 1},
		{name: "smoke 2", args: args{nums: []int{3, 1, 4, 3, 2, 2, 4}, k: 2}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countGood(test.args.nums, test.args.k); got != test.want {
				t.Errorf("countGood() = %v, want = %v", got, test.want)
			}
		})
	}
}
