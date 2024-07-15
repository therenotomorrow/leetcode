package golang

import "testing"

func TestMinimumOperations2(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 4}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{3, 6, 9}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumOperations2(test.args.nums); got != test.want {
				t.Errorf("minimumOperations2() = %v, want = %v", got, test.want)
			}
		})
	}
}
