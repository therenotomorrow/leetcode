package golang

import "testing"

func TestMaxWidthOfVerticalArea(t *testing.T) {
	t.Parallel()

	type args struct {
		points [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{points: [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}}, want: 1},
		{name: "smoke 2", args: args{points: [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxWidthOfVerticalArea(test.args.points); got != test.want {
				t.Errorf("maxWidthOfVerticalArea() = %v, want = %v", got, test.want)
			}
		})
	}
}
