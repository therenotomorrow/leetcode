package golang

import "testing"

func TestCountComponents(t *testing.T) {
	t.Parallel()

	type args struct {
		n     int
		edges [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 5, edges: [][]int{{0, 1}, {1, 2}, {3, 4}}}, want: 2},
		{name: "smoke 2", args: args{n: 5, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countComponents(test.args.n, test.args.edges); got != test.want {
				t.Errorf("countComponents() = %v, want = %v", got, test.want)
			}
		})
	}
}
