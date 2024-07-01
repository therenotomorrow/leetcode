package golang

import "testing"

func TestFindCenter(t *testing.T) {
	t.Parallel()

	type args struct {
		edges [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{edges: [][]int{{1, 2}, {2, 3}, {4, 2}}}, want: 2},
		{name: "smoke 2", args: args{edges: [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findCenter(test.args.edges); got != test.want {
				t.Errorf("findCenter() = %v, want = %v", got, test.want)
			}
		})
	}
}
