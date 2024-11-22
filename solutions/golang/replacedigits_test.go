package golang

import (
	"testing"
)

func TestReplaceDigits(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "a1c1e1"}, want: "abcdef"},
		{name: "smoke 2", args: args{s: "a1b2c3d4e"}, want: "abbdcfdhe"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := replaceDigits(test.args.s); got != test.want {
				t.Errorf("replaceDigits() = %v, want = %v", got, test.want)
			}
		})
	}
}
