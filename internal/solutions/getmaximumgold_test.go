package solutions

import "testing"

func TestGetMaximumGold(t *testing.T) {
	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}}, want: 24},
		{name: "smoke 2", args: args{grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}}, want: 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGold(tt.args.grid); got != tt.want {
				t.Errorf("getMaximumGold() = %v, want = %v", got, tt.want)
			}
		})
	}
}
