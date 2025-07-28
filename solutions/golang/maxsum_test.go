package golang

import "testing"

func TestMaxSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 4, 5}}, want: 15},
		{name: "smoke 2", args: args{nums: []int{1, 1, 0, 1, 1}}, want: 1},
		{name: "smoke 3", args: args{nums: []int{1, 2, -1, -2, 1, 0, -1}}, want: 3},
		{name: "test 808: wrong answer", args: args{nums: []int{-100}}, want: -100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxSum(test.args.nums); got != test.want {
				t.Errorf("maxSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
