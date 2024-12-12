package golang

import (
	"testing"
)

func TestMaximumLength(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aaaa"}, want: 2},
		{name: "smoke 2", args: args{s: "abcdef"}, want: -1},
		{name: "smoke 3", args: args{s: "abcaba"}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumLength(test.args.s); got != test.want {
				t.Errorf("maximumLength() = %v, want = %v", got, test.want)
			}
		})
	}
}
