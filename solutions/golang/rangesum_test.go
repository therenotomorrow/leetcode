package golang

import "testing"

func TestRangeSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums  []int
		n     int
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 4}, n: 4, left: 1, right: 5}, want: 13},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4}, n: 4, left: 3, right: 4}, want: 6},
		{name: "smoke 3", args: args{nums: []int{1, 2, 3, 4}, n: 4, left: 1, right: 10}, want: 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := rangeSum(test.args.nums, test.args.n, test.args.left, test.args.right); got != test.want {
				t.Errorf("rangeSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
