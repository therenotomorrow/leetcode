package golang

import "testing"

func TestMinDistance(t *testing.T) {
	t.Parallel()

	type args struct {
		height   int
		width    int
		tree     []int
		squirrel []int
		nuts     [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				height:   5,
				width:    6,
				tree:     []int{2, 2},
				squirrel: []int{4, 4},
				nuts:     [][]int{{3, 0}, {2, 5}},
			},
			want: 12,
		},
		{
			name: "smoke 2",
			args: args{
				height:   1,
				width:    3,
				tree:     []int{0, 1},
				squirrel: []int{0, 0},
				nuts:     [][]int{{0, 2}},
			},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := minDistance(test.args.height, test.args.width, test.args.tree, test.args.squirrel, test.args.nuts)

			if got != test.want {
				t.Errorf("minDistance() = %v, want = %v", got, test.want)
			}
		})
	}
}
