package golang

import (
	"testing"
)

func TestFindLengthOfShortestSubarray(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{1, 2, 3, 10, 4, 2, 3, 5}}, want: 3},
		{name: "smoke 2", args: args{arr: []int{5, 4, 3, 2, 1}}, want: 4},
		{name: "smoke 3", args: args{arr: []int{1, 2, 3}}, want: 0},
		{name: "test 78: wrong answer", args: args{arr: []int{2, 2, 2, 1, 1, 1}}, want: 3},
		{name: "test 82: wrong answer", args: args{arr: []int{6, 3, 10, 11, 15, 20, 13, 3, 18, 12}}, want: 8},
		{name: "test 115: wrong answer", args: args{arr: []int{1, 2, 3, 3, 10, 1, 3, 3, 5}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findLengthOfShortestSubarray(test.args.arr); got != test.want {
				t.Errorf("findLengthOfShortestSubarray() = %v, want = %v", got, test.want)
			}
		})
	}
}
