package golang

import "testing"

func TestBuyChoco(t *testing.T) {
	t.Parallel()

	type args struct {
		prices []int
		money  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{prices: []int{1, 2, 2}, money: 3}, want: 0},
		{name: "smoke 2", args: args{prices: []int{3, 2, 3}, money: 3}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := buyChoco(test.args.prices, test.args.money); got != test.want {
				t.Errorf("buyChoco() = %v, want = %v", got, test.want)
			}
		})
	}
}
