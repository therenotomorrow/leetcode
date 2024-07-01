package golang

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 0, 3, 12}}, want: []int{1, 3, 12, 0, 0}},
		{name: "smoke 2", args: args{nums: []int{0}}, want: []int{0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			moveZeroes(test.args.nums)

			if !reflect.DeepEqual(test.args.nums, test.want) {
				t.Errorf("moveZeroes() = %v, want = %v", test.args.nums, test.want)
			}
		})
	}
}
