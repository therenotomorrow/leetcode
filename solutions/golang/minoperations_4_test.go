package golang

import (
	"testing"
)

func TestMinOperations4(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 1, 1}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 5, 2, 4, 1}}, want: 14},
		{name: "smoke 3", args: args{nums: []int{8}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations4(test.args.nums); got != test.want {
				t.Errorf("minOperations4() = %v, want = %v", got, test.want)
			}
		})
	}
}
