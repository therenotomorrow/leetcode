package solutions

import "testing"

func TestBuyChoco(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buyChoco(tt.args.prices, tt.args.money); got != tt.want {
				t.Errorf("buyChoco() = %v, want = %v", got, tt.want)
			}
		})
	}
}
