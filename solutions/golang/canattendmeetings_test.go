package golang

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	t.Parallel()

	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{intervals: [][]int{{0, 30}, {5, 10}, {15, 20}}}, want: false},
		{name: "smoke 2", args: args{intervals: [][]int{{7, 10}, {2, 4}}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canAttendMeetings(test.args.intervals); got != test.want {
				t.Errorf("canAttendMeetings() = %v, want = %v", got, test.want)
			}
		})
	}
}
