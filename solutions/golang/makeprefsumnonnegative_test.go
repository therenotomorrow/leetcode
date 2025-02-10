package golang

import "testing"

func TestMakePrefSumNonNegative(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 3, -5, 4}}, want: 0},
		{name: "smoke 2", args: args{nums: []int{3, -5, -2, 6}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := makePrefSumNonNegative(test.args.nums); got != test.want {
				t.Errorf("makePrefSumNonNegative() = %v, want = %v", got, test.want)
			}
		})
	}
}
