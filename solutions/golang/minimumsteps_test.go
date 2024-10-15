package golang

import "testing"

func TestMinimumSteps(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{s: "101"}, want: 1},
		{name: "smoke 2", args: args{s: "100"}, want: 2},
		{name: "smoke 3", args: args{s: "0111"}, want: 0},
		{name: "own 1", args: args{s: "1110"}, want: 3},
		{name: "test 97: wrong answer", args: args{s: "11000111"}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumSteps(test.args.s); got != test.want {
				t.Errorf("minimumSteps() = %v, want = %v", got, test.want)
			}
		})
	}
}
