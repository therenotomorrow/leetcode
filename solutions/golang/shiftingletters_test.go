package golang

import "testing"

func TestShiftingLetters(t *testing.T) {
	t.Parallel()

	type args struct {
		s      string
		shifts [][]int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "abc", shifts: [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}}, want: "ace"},
		{name: "smoke 2", args: args{s: "dztz", shifts: [][]int{{0, 0, 0}, {1, 1, 1}}}, want: "catz"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := shiftingLetters(test.args.s, test.args.shifts); got != test.want {
				t.Errorf("shiftingLetters() = %v, want = %v", got, test.want)
			}
		})
	}
}
