package golang

import "testing"

func TestFindMinDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		timePoints []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{timePoints: []string{"23:59", "00:00"}}, want: 1},
		{name: "smoke 2", args: args{timePoints: []string{"00:00", "23:59", "00:00"}}, want: 0},
		{name: "test 32: wrong answer", args: args{timePoints: []string{"01:01", "02:01", "03:00"}}, want: 59},
		{name: "test 109: wrong answer", args: args{timePoints: []string{"00:00", "04:00", "22:00"}}, want: 120},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMinDifference(test.args.timePoints); got != test.want {
				t.Errorf("findMinDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
