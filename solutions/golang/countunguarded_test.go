package golang

import (
	"testing"
)

func TestCountUnguarded(t *testing.T) {
	t.Parallel()

	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{m: 4, n: 6, guards: [][]int{{0, 0}, {1, 1}, {2, 3}}, walls: [][]int{{0, 1}, {2, 2}, {1, 4}}},
			want: 7,
		},
		{
			name: "smoke 2",
			args: args{m: 3, n: 3, guards: [][]int{{1, 1}}, walls: [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}},
			want: 4,
		},
		{
			name: "test 13: wrong answer",
			args: args{m: 5, n: 5, guards: [][]int{{1, 4}, {4, 1}, {0, 3}}, walls: [][]int{{3, 2}}},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countUnguarded(test.args.m, test.args.n, test.args.guards, test.args.walls); got != test.want {
				t.Errorf("countUnguarded() = %v, want = %v", got, test.want)
			}
		})
	}
}
