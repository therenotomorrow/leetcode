package golang

import "testing"

func TestMinOperations7(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 1, 1, 0, 0}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{0, 1, 1, 1}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations7(test.args.nums); got != test.want {
				t.Errorf("minOperations7() = %v, want = %v", got, test.want)
			}
		})
	}
}
