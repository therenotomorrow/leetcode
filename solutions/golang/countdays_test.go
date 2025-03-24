package golang

import "testing"

func TestCountDays(t *testing.T) {
	t.Parallel()

	type args struct {
		days     int
		meetings [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{days: 10, meetings: [][]int{{5, 7}, {1, 3}, {9, 10}}}, want: 2},
		{name: "smoke 2", args: args{days: 5, meetings: [][]int{{2, 4}, {1, 3}}}, want: 1},
		{name: "smoke 3", args: args{days: 6, meetings: [][]int{{1, 6}}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countDays(test.args.days, test.args.meetings); got != test.want {
				t.Errorf("countDays() = %v, want = %v", got, test.want)
			}
		})
	}
}
