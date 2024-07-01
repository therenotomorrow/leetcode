package golang

import "testing"

func TestSumOfDigits(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{34, 23, 1, 24, 75, 33, 54, 8}}, want: 0},
		{name: "smoke 2", args: args{nums: []int{99, 77, 33, 66, 55}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sumOfDigits(test.args.nums); got != test.want {
				t.Errorf("sumOfDigits() = %v, want = %v", got, test.want)
			}
		})
	}
}
