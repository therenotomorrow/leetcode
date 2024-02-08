package solutions

import "testing"

func TestNumSquares(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 12}, want: 3},
		{name: "smoke 2", args: args{n: 13}, want: 2},
		{name: "test 587: wrong answer", args: args{n: 1}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want = %v", got, tt.want)
			}
		})
	}
}
