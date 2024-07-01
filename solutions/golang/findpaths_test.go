package golang

import "testing"

func TestFindPaths(t *testing.T) {
	t.Parallel()

	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{m: 2, n: 2, maxMove: 2, startRow: 0, startColumn: 0}, want: 6},
		{name: "smoke 2", args: args{m: 1, n: 3, maxMove: 3, startRow: 0, startColumn: 1}, want: 12},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := findPaths(test.args.m, test.args.n, test.args.maxMove, test.args.startRow, test.args.startColumn)

			if got != test.want {
				t.Errorf("findPaths() = %v, want = %v", got, test.want)
			}
		})
	}
}
