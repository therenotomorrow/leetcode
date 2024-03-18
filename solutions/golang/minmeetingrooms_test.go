package golang

import "testing"

func TestMinMeetingRooms(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want = %v", got, tt.want)
			}
		})
	}
}
