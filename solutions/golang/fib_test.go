package golang

import "testing"

func TestFib(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 2}, want: 1},
		{name: "smoke 2", args: args{n: 3}, want: 2},
		{name: "smoke 3", args: args{n: 4}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := fib(test.args.n); got != test.want {
				t.Errorf("fib() = %v, want = %v", got, test.want)
			}
		})
	}
}
