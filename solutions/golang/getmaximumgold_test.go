package golang

import "testing"

func TestGetMaximumGold(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getMaximumGold(test.args.grid); got != test.want {
				t.Errorf("getMaximumGold() = %v, want = %v", got, test.want)
			}
		})
	}
}
