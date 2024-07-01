package golang

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minTimeToVisitAllPoints(test.args.points); got != test.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want = %v", got, test.want)
			}
		})
	}
}
