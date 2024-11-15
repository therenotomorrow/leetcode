package golang

import (
	"testing"
)

func TestGetMaximumGenerated(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 7}, want: 3},
		{name: "smoke 2", args: args{n: 2}, want: 1},
		{name: "smoke 3", args: args{n: 3}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getMaximumGenerated(test.args.n); got != test.want {
				t.Errorf("getMaximumGenerated() = %v, want = %v", got, test.want)
			}
		})
	}
}
