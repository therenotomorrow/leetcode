package golang

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{name: "smoke 2", args: args{nums: []int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sortedSquares(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("sortedSquares() = %v, want = %v", got, test.want)
			}
		})
	}
}
