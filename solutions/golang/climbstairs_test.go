package golang

import "testing"

func TestClimbStairs(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want = %v", got, tt.want)
			}
		})
	}
}
