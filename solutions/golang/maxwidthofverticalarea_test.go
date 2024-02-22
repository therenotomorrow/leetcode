package golang

import "testing"

func TestMaxWidthOfVerticalArea(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthOfVerticalArea(tt.args.points); got != tt.want {
				t.Errorf("maxWidthOfVerticalArea() = %v, want = %v", got, tt.want)
			}
		})
	}
}
