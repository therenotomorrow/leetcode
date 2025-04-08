package golang

import "testing"

func TestMinimumOperations3(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{4, 5, 6, 4, 4}}, want: 2},
		{name: "smoke 3", args: args{nums: []int{6, 7, 8, 9}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumOperations3(test.args.nums); got != test.want {
				t.Errorf("minimumOperations3() = %v, want = %v", got, test.want)
			}
		})
	}
}
