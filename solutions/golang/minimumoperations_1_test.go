package golang

import "testing"

func TestMinimumOperations1(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 5, 0, 3, 5}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{0}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumOperations1(test.args.nums); got != test.want {
				t.Errorf("minimumOperations1() = %v, want = %v", got, test.want)
			}
		})
	}
}
