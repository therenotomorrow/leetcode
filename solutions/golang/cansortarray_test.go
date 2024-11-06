package golang

import (
	"testing"
)

func TestCanSortArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{8, 4, 2, 30, 15}}, want: true},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4, 5}}, want: true},
		{name: "smoke 3", args: args{nums: []int{3, 16, 8, 4, 2}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canSortArray(test.args.nums); got != test.want {
				t.Errorf("canSortArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
