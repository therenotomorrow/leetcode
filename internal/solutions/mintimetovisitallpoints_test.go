package solutions

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{points: [][]int{{1, 1}, {3, 4}, {-1, 0}}}, want: 7},
		{name: "smoke 2", args: args{points: [][]int{{3, 2}, {-2, 2}}}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want = %v", got, tt.want)
			}
		})
	}
}
