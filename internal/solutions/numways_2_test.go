package solutions

import "testing"

func TestNumWays2(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays2(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("numWays2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
