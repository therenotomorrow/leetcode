package golang

import "testing"

func TestMinBitFlips(t *testing.T) {
	t.Parallel()

	type args struct {
		start int
		goal  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{start: 10, goal: 7}, want: 3},
		{name: "smoke 2", args: args{start: 3, goal: 4}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minBitFlips(test.args.start, test.args.goal); got != test.want {
				t.Errorf("minBitFlips() = %v, want = %v", got, test.want)
			}
		})
	}
}
