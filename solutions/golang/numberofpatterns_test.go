package golang

import "testing"

func TestNumberOfPatterns(t *testing.T) {
	t.Parallel()

	type args struct {
		m int
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{m: 1, n: 1}, want: 9},
		{name: "smoke 2", args: args{m: 1, n: 2}, want: 65},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfPatterns(test.args.m, test.args.n); got != test.want {
				t.Errorf("numberOfPatterns() = %v, want = %v", got, test.want)
			}
		})
	}
}
