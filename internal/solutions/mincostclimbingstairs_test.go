package solutions

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{cost: []int{10, 15, 20}}, want: 15},
		{name: "smoke 2", args: args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want = %v", got, tt.want)
			}
		})
	}
}
