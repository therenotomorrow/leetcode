package golang

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	t.Parallel()

	type args struct {
		points [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, want: 2},
		{name: "smoke 2", args: args{points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}}, want: 4},
		{name: "smoke 3", args: args{points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMinArrowShots(test.args.points); got != test.want {
				t.Errorf("findMinArrowShots() = %v, want = %v", got, test.want)
			}
		})
	}
}
