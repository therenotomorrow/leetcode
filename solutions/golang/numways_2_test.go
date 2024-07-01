package golang

import "testing"

func TestNumWays2(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 3, k: 2}, want: 6},
		{name: "smoke 2", args: args{n: 1, k: 1}, want: 1},
		{name: "smoke 3", args: args{n: 7, k: 2}, want: 42},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numWays2(test.args.n, test.args.k); got != test.want {
				t.Errorf("numWays2() = %v, want = %v", got, test.want)
			}
		})
	}
}
