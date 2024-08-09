package golang

import "testing"

func TestNumMagicSquaresInside(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{grid: [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}}, want: 1},
		{name: "smoke 2", args: args{grid: [][]int{{8}}}, want: 0},
		{name: "test 92: wrong answer", args: args{grid: [][]int{{2, 7, 6}, {1, 5, 9}, {4, 3, 8}}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numMagicSquaresInside(test.args.grid); got != test.want {
				t.Errorf("numMagicSquaresInside() = %v, want = %v", got, test.want)
			}
		})
	}
}
