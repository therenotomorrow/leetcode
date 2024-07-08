package golang

import "testing"

func TestFindTheWinner(t *testing.T) {
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
		{name: "smoke 1", args: args{n: 5, k: 2}, want: 3},
		{name: "smoke 2", args: args{n: 6, k: 5}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findTheWinner(test.args.n, test.args.k); got != test.want {
				t.Errorf("findTheWinner() = %v, want = %v", got, test.want)
			}
		})
	}
}
