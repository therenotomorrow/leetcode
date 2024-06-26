package golang

import "testing"

func TestTribonacci(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 4}, want: 4},
		{name: "smoke 2", args: args{n: 25}, want: 1389537},
		{name: "own 1", args: args{n: 0}, want: 0},
		{name: "own 2", args: args{n: 1}, want: 1},
		{name: "own 3", args: args{n: 2}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := tribonacci(test.args.n); got != test.want {
				t.Errorf("tribonacci() = %v, want = %v", got, test.want)
			}
		})
	}
}
