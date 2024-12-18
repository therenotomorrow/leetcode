package golang

import (
	"testing"
)

func TestRepeatLimitedString(t *testing.T) {
	t.Parallel()

	type args struct {
		s           string
		repeatLimit int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "cczazcc", repeatLimit: 3}, want: "zzcccac"},
		{name: "smoke 2", args: args{s: "aababab", repeatLimit: 2}, want: "bbabaa"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := repeatLimitedString(test.args.s, test.args.repeatLimit); got != test.want {
				t.Errorf("repeatLimitedString() = %v, want = %v", got, test.want)
			}
		})
	}
}
