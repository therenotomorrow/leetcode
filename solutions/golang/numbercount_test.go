package golang

import (
	"testing"
)

func TestNumberCount(t *testing.T) {
	t.Parallel()

	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{a: 1, b: 20}, want: 19},
		{name: "smoke 2", args: args{a: 9, b: 19}, want: 10},
		{name: "smoke 3", args: args{a: 80, b: 120}, want: 27},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberCount(test.args.a, test.args.b); got != test.want {
				t.Errorf("numberCount() = %v, want = %v", got, test.want)
			}
		})
	}
}
