package golang

import (
	"testing"
)

func TestCheck(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{3, 4, 5, 1, 2}}, want: true},
		{name: "smoke 2", args: args{nums: []int{2, 1, 3, 4}}, want: false},
		{name: "smoke 3", args: args{nums: []int{1, 2, 3}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := check(test.args.nums); got != test.want {
				t.Errorf("check() = %v, want = %v", got, test.want)
			}
		})
	}
}
