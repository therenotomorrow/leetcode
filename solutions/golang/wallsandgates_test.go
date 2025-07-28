package golang

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	t.Parallel()

	type args struct {
		rooms [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{rooms: [][]int{
				{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647},
			}},
			want: [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}},
		},
		{
			name: "smoke 2",
			args: args{rooms: [][]int{{-1}}},
			want: [][]int{{-1}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			wallsAndGates(test.args.rooms)

			if !reflect.DeepEqual(test.args.rooms, test.want) {
				t.Errorf("wallsAndGates() = %v, want = %v", test.args.rooms, test.want)
			}
		})
	}
}
