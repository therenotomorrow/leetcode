package golang

import (
	"testing"
)

func TestMaxMatrixSum(t *testing.T) {
	t.Parallel()

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{matrix: [][]int{{1, -1}, {-1, 1}}}, want: 4},
		{name: "smoke 2", args: args{matrix: [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}}, want: 16},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxMatrixSum(test.args.matrix); got != test.want {
				t.Errorf("maxMatrixSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
