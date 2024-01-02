package solutions

import (
	"reflect"
	"testing"
)

func TestFindMissingRanges(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingRanges(tt.args.nums, tt.args.lower, tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want = %v", got, tt.want)
			}
		})
	}
}
