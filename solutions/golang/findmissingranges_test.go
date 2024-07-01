package golang

import (
	"reflect"
	"testing"
)

func TestFindMissingRanges(t *testing.T) {
	t.Parallel()

	type args struct {
		nums  []int
		lower int
		upper int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{nums: []int{0, 1, 3, 50, 75}, lower: 0, upper: 99},
			want: [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}},
		},
		{
			name: "smoke 2",
			args: args{nums: []int{-1}, lower: -1, upper: -1},
			want: [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMissingRanges(test.args.nums, test.args.lower, test.args.upper); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findMissingRanges() = %v, want = %v", got, test.want)
			}
		})
	}
}
