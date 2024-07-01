package golang

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	t.Parallel()

	type args struct {
		n      int
		k      int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 1, k: 6, target: 3}, want: 1},
		{name: "smoke 2", args: args{n: 2, k: 6, target: 7}, want: 6},
		{name: "smoke 3", args: args{n: 30, k: 30, target: 500}, want: 222616187},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numRollsToTarget(test.args.n, test.args.k, test.args.target); got != test.want {
				t.Errorf("numRollsToTarget() = %v, want = %v", got, test.want)
			}
		})
	}
}
