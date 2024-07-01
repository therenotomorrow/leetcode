package golang

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 4}}, want: []int{24, 12, 8, 6}},
		{name: "smoke 2", args: args{nums: []int{-1, 1, 0, -3, 3}}, want: []int{0, 0, 9, 0, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := productExceptSelf(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("productExceptSelf() = %v, want = %v", got, test.want)
			}
		})
	}
}
