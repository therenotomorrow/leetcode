package golang

import "testing"

func TestMaximumSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{18, 43, 36, 13, 7}}, want: 54},
		{name: "smoke 2", args: args{nums: []int{10, 12, 19, 14}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumSum(test.args.nums); got != test.want {
				t.Errorf("maximumSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
