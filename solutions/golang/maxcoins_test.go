package golang

import "testing"

func TestMaxCoins(t *testing.T) {
	t.Parallel()

	type args struct {
		piles []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{piles: []int{2, 4, 1, 2, 7, 8}}, want: 9},
		{name: "smoke 2", args: args{piles: []int{2, 4, 5}}, want: 4},
		{name: "smoke 3", args: args{piles: []int{9, 8, 7, 6, 5, 1, 2, 3, 4}}, want: 18},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxCoins(test.args.piles); got != test.want {
				t.Errorf("maxCoins() = %v, want = %v", got, test.want)
			}
		})
	}
}
