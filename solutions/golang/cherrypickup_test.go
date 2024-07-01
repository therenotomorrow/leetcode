package golang

import "testing"

func TestCherryPickup(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{grid: [][]int{
				{3, 1, 1},
				{2, 5, 1},
				{1, 5, 5},
				{2, 1, 1},
			}},
			want: 24,
		},
		{
			name: "smoke 2",
			args: args{grid: [][]int{
				{1, 0, 0, 0, 0, 0, 1},
				{2, 0, 0, 0, 0, 3, 0},
				{2, 0, 9, 0, 0, 0, 0},
				{0, 3, 0, 5, 4, 0, 0},
				{1, 0, 2, 3, 0, 0, 6},
			}},
			want: 28,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := cherryPickup(test.args.grid); got != test.want {
				t.Errorf("cherryPickup() = %v, want = %v", got, test.want)
			}
		})
	}
}
