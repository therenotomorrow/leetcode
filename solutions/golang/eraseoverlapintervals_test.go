package golang

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	t.Parallel()

	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, want: 1},
		{name: "smoke 2", args: args{intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}}, want: 2},
		{name: "smoke 3", args: args{intervals: [][]int{{1, 2}, {2, 3}}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := eraseOverlapIntervals(test.args.intervals); got != test.want {
				t.Errorf("eraseOverlapIntervals() = %v, want = %v", got, test.want)
			}
		})
	}
}
