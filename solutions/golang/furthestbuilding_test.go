package golang

import "testing"

func TestFurthestBuilding(t *testing.T) {
	t.Parallel()

	type args struct {
		heights []int
		bricks  int
		ladders int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{heights: []int{4, 2, 7, 6, 9, 14, 12}, bricks: 5, ladders: 1},
			want: 4,
		},
		{
			name: "smoke 2",
			args: args{heights: []int{4, 12, 2, 7, 3, 18, 20, 3, 19}, bricks: 10, ladders: 2},
			want: 7,
		},
		{
			name: "smoke 3",
			args: args{heights: []int{14, 3, 19, 3}, bricks: 17, ladders: 0},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := furthestBuilding(test.args.heights, test.args.bricks, test.args.ladders); got != test.want {
				t.Errorf("furthestBuilding() = %v, want = %v", got, test.want)
			}
		})
	}
}
