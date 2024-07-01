package golang

import "testing"

func TestMinTotalDistance(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minTotalDistance(test.args.grid); got != test.want {
				t.Errorf("minTotalDistance() = %v, want = %v", got, test.want)
			}
		})
	}
}
