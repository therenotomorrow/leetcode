package golang

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 2}, want: 2},
		{name: "smoke 2", args: args{n: 3}, want: 3},
		{name: "test 2: runtime", args: args{n: 1}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := climbStairs(test.args.n); got != test.want {
				t.Errorf("climbStairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
