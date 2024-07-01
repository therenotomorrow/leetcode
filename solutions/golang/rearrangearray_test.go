package golang

import (
	"reflect"
	"testing"
)

func TestRearrangeArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{3, 1, -2, -5, 2, -4}}, want: []int{3, -2, 1, -5, 2, -4}},
		{name: "smoke 2", args: args{nums: []int{-1, 1}}, want: []int{1, -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := rearrangeArray(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("rearrangeArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
