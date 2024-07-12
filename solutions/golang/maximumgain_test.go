package golang

import "testing"

func TestMaximumGain(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "cdbcbbaaabab", x: 4, y: 5}, want: 19},
		{name: "smoke 2", args: args{s: "aabbaaxybbaabb", x: 5, y: 4}, want: 20},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumGain(test.args.s, test.args.x, test.args.y); got != test.want {
				t.Errorf("maximumGain() = %v, want = %v", got, test.want)
			}
		})
	}
}
