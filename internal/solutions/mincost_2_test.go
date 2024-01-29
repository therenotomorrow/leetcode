package solutions

import "testing"

func TestMinCost2(t *testing.T) {
	type args struct {
		costs [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{costs: [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}}, want: 10},
		{name: "smoke 2", args: args{costs: [][]int{{7, 6, 2}}}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost2(tt.args.costs); got != tt.want {
				t.Errorf("minCost2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
