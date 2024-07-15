package golang

import "testing"

func TestDifferenceOfSums(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
		m int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 10, m: 3}, want: 19},
		{name: "smoke 2", args: args{n: 5, m: 6}, want: 15},
		{name: "smoke 3", args: args{n: 5, m: 1}, want: -15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := differenceOfSums(test.args.n, test.args.m); got != test.want {
				t.Errorf("differenceOfSums() = %v, want = %v", got, test.want)
			}
		})
	}
}
