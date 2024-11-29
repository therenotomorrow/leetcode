package golang

import (
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	t.Parallel()

	type args struct {
		rooms [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{rooms: [][]int{{1}, {2}, {3}, {}}}, want: true},
		{name: "smoke 2", args: args{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canVisitAllRooms(test.args.rooms); got != test.want {
				t.Errorf("canVisitAllRooms() = %v, want = %v", got, test.want)
			}
		})
	}
}
