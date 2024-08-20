package golang

import "testing"

func TestArrangeCoins(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 5}, want: 2},
		{name: "smoke 2", args: args{n: 8}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := arrangeCoins(test.args.n); got != test.want {
				t.Errorf("arrangeCoins() = %v, want = %v", got, test.want)
			}
		})
	}
}
