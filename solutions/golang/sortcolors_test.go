package golang

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 0, 2, 1, 1, 0}}, want: []int{0, 0, 1, 1, 2, 2}},
		{name: "smoke 2", args: args{nums: []int{2, 0, 1}}, want: []int{0, 1, 2}},
		{name: "test 46: wrong answer", args: args{nums: []int{1, 2, 0}}, want: []int{0, 1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			sortColors(test.args.nums)

			if !reflect.DeepEqual(test.args.nums, test.want) {
				t.Errorf("sortColors() = %v, want = %v", test.args.nums, test.want)
			}
		})
	}
}
