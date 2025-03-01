package golang

import "testing"

func TestMaxAbsoluteSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, -3, 2, 3, -4}}, want: 5},
		{name: "smoke 2", args: args{nums: []int{2, -5, 1, -4, 3, -2}}, want: 8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxAbsoluteSum(test.args.nums); got != test.want {
				t.Errorf("maxAbsoluteSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
