package golang

import (
	"testing"
)

func TestMaxEqualRowsAfterFlips(t *testing.T) {
	t.Parallel()

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{matrix: [][]int{{0, 1}, {1, 1}}}, want: 1},
		{name: "smoke 2", args: args{matrix: [][]int{{0, 1}, {1, 0}}}, want: 2},
		{name: "smoke 3", args: args{matrix: [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxEqualRowsAfterFlips(test.args.matrix); got != test.want {
				t.Errorf("maxEqualRowsAfterFlips() = %v, want = %v", got, test.want)
			}
		})
	}
}
