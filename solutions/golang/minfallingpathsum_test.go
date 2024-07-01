package golang

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	t.Parallel()

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}}, want: 13},
		{name: "smoke 2", args: args{matrix: [][]int{{-19, 57}, {-40, -5}}}, want: -59},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minFallingPathSum(test.args.matrix); got != test.want {
				t.Errorf("minFallingPathSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
