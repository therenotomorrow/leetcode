package golang

import (
	"reflect"
	"testing"
)

func TestSpiralMatrixIII(t *testing.T) {
	t.Parallel()

	type args struct {
		rows   int
		cols   int
		rStart int
		cStart int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1", args: args{rows: 1, cols: 4, rStart: 0, cStart: 0}, want: [][]int{
				{0, 0}, {0, 1}, {0, 2}, {0, 3},
			},
		},
		{
			name: "smoke 2", args: args{rows: 5, cols: 6, rStart: 1, cStart: 4}, want: [][]int{
				{1, 4},
				{1, 5},
				{2, 5},
				{2, 4},
				{2, 3},
				{1, 3},
				{0, 3},
				{0, 4},
				{0, 5},
				{3, 5},
				{3, 4},
				{3, 3},
				{3, 2},
				{2, 2},
				{1, 2},
				{0, 2},
				{4, 5},
				{4, 4},
				{4, 3},
				{4, 2},
				{4, 1},
				{3, 1},
				{2, 1},
				{1, 1},
				{0, 1},
				{4, 0},
				{3, 0},
				{2, 0},
				{1, 0},
				{0, 0},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := spiralMatrixIII(test.args.rows, test.args.cols, test.args.rStart, test.args.cStart)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("spiralMatrixIII() = %v, want = %v", got, test.want)
			}
		})
	}
}
