package golang

import (
	"testing"
)

func TestMinEnd(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
		x int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{n: 3, x: 4}, want: 6},
		{name: "smoke 2", args: args{n: 2, x: 7}, want: 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minEnd(test.args.n, test.args.x); got != test.want {
				t.Errorf("minEnd() = %v, want = %v", got, test.want)
			}
		})
	}
}
