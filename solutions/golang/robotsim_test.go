package golang

import "testing"

func TestRobotSim(t *testing.T) {
	t.Parallel()

	type args struct {
		commands  []int
		obstacles [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{commands: []int{4, -1, 3}, obstacles: [][]int{}}, want: 25},
		{name: "smoke 2", args: args{commands: []int{4, -1, 4, -2, 4}, obstacles: [][]int{{2, 4}}}, want: 65},
		{name: "smoke 3", args: args{commands: []int{6, -1, -1, 6}, obstacles: [][]int{}}, want: 36},
		{
			name: "test 6, wrong answer",
			args: args{
				commands:  []int{-2, -1, 4, 7, 8},
				obstacles: [][]int{{1, 1}, {2, 1}, {4, 4}, {5, -5}, {2, -3}, {-2, -3}, {-1, -3}, {-4, -1}, {-4, 3}, {5, 1}},
			},
			want: 361,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := robotSim(test.args.commands, test.args.obstacles); got != test.want {
				t.Errorf("robotSim() = %v, want = %v", got, test.want)
			}
		})
	}
}
