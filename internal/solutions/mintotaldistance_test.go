package solutions

import "testing"

func TestMinTotalDistance(t *testing.T) {
	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{grid: [][]int{{1, 0, 0, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}}, want: 6},
		{name: "smoke 2", args: args{grid: [][]int{{1, 1}}}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTotalDistance(tt.args.grid); got != tt.want {
				t.Errorf("minTotalDistance() = %v, want = %v", got, tt.want)
			}
		})
	}
}
