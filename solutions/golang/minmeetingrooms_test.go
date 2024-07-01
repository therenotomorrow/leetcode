package golang

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	t.Parallel()

	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{intervals: [][]int{{0, 30}, {5, 10}, {15, 20}}}, want: 2},
		{name: "smoke 2", args: args{intervals: [][]int{{7, 10}, {2, 4}}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minMeetingRooms(test.args.intervals); got != test.want {
				t.Errorf("minMeetingRooms() = %v, want = %v", got, test.want)
			}
		})
	}
}
