package golang

import "testing"

func TestMinCost2(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minCost2(test.args.costs); got != test.want {
				t.Errorf("minCost2() = %v, want = %v", got, test.want)
			}
		})
	}
}
