package golang

import "testing"

func TestMinMoves(t *testing.T) {
	t.Parallel()

	type args struct {
		rooks [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{rooks: [][]int{{0, 0}, {1, 0}, {1, 1}}}, want: 3},
		{name: "smoke 2", args: args{rooks: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minMoves(test.args.rooks); got != test.want {
				t.Errorf("minMoves() = %v, want = %v", got, test.want)
			}
		})
	}
}
