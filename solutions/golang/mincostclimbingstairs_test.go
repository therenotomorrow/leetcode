package golang

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minCostClimbingStairs(test.args.cost); got != test.want {
				t.Errorf("minCostClimbingStairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
