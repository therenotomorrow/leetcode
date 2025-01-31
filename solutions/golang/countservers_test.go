package golang

import "testing"

func TestCountServers(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{grid: [][]int{{1, 0}, {0, 1}}}, want: 0},
		{name: "smoke 2", args: args{grid: [][]int{{1, 0}, {1, 1}}}, want: 3},
		{name: "smoke 3", args: args{grid: [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countServers(test.args.grid); got != test.want {
				t.Errorf("countServers() = %v, want = %v", got, test.want)
			}
		})
	}
}
