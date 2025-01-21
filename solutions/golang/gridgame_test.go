package golang

import "testing"

func TestGridGame(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{grid: [][]int{{2, 5, 4}, {1, 5, 1}}}, want: 4},
		{name: "smoke 2", args: args{grid: [][]int{{3, 3, 1}, {8, 5, 2}}}, want: 4},
		{name: "smoke 3", args: args{grid: [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}}, want: 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := gridGame(test.args.grid); got != test.want {
				t.Errorf("gridGame() = %v, want = %v", got, test.want)
			}
		})
	}
}
