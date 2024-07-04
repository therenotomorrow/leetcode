package golang

import "testing"

func TestEarliestAcq(t *testing.T) {
	t.Parallel()

	type args struct {
		logs [][]int
		n    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{logs: [][]int{
			{20190101, 0, 1},
			{20190104, 3, 4},
			{20190107, 2, 3},
			{20190211, 1, 5},
			{20190224, 2, 4},
			{20190301, 0, 3},
			{20190312, 1, 2},
			{20190322, 4, 5},
		}, n: 6}, want: 20190301},
		{name: "smoke 2", args: args{logs: [][]int{
			{0, 2, 0},
			{1, 0, 1},
			{3, 0, 3},
			{4, 1, 2},
			{7, 3, 1},
		}, n: 4}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := earliestAcq(test.args.logs, test.args.n); got != test.want {
				t.Errorf("earliestAcq() = %v, want = %v", got, test.want)
			}
		})
	}
}
