package golang

import "testing"

func TestColoredCells(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{n: 1}, want: 1},
		{name: "smoke 2", args: args{n: 2}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := coloredCells(test.args.n); got != test.want {
				t.Errorf("coloredCells() = %v, want = %v", got, test.want)
			}
		})
	}
}
