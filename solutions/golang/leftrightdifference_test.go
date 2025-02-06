package golang

import (
	"reflect"
	"testing"
)

func TestLeftRightDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{10, 4, 8, 3}}, want: []int{15, 1, 11, 22}},
		{name: "smoke 2", args: args{nums: []int{1}}, want: []int{0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := leftRightDifference(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("leftRightDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
