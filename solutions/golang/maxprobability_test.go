package golang

import "testing"

func TestMaxProbability(t *testing.T) {
	t.Parallel()

	type args struct {
		n         int
		edges     [][]int
		probs     []float64
		startNode int
		endNode   int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "smoke 1",
			args: args{
				n:         3,
				edges:     [][]int{{0, 1}, {1, 2}, {0, 2}},
				probs:     []float64{0.5, 0.5, 0.2},
				startNode: 0,
				endNode:   2,
			},
			want: 0.25,
		},
		{
			name: "smoke 2",
			args: args{
				n:         3,
				edges:     [][]int{{0, 1}, {1, 2}, {0, 2}},
				probs:     []float64{0.5, 0.5, 0.3},
				startNode: 0,
				endNode:   2,
			},
			want: 0.3,
		},
		{
			name: "smoke 3",
			args: args{
				n:         3,
				edges:     [][]int{{0, 1}},
				probs:     []float64{0.5},
				startNode: 0,
				endNode:   2,
			},
			want: 0,
		},
		{
			name: "test 8: wrong answer",
			args: args{
				n:         5,
				edges:     [][]int{{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3}},
				probs:     []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04},
				startNode: 3,
				endNode:   4,
			},
			want: 0.2139,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := maxProbability(
				test.args.n,
				test.args.edges,
				test.args.probs,
				test.args.startNode,
				test.args.endNode,
			)

			if got != test.want {
				t.Errorf("maxProbability() = %v, want = %v", got, test.want)
			}
		})
	}
}
