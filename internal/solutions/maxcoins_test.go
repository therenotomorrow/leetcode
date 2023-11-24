package solutions

import "testing"

func TestMaxCoins(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.piles); got != tt.want {
				t.Errorf("maxCoins() = %v, want = %v", got, tt.want)
			}
		})
	}
}
