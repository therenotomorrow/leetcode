package golang

import (
	"testing"
)

func TestTakeCharacters(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aabaaaacaabc", k: 2}, want: 8},
		{name: "smoke 2", args: args{s: "a", k: 1}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := takeCharacters(test.args.s, test.args.k); got != test.want {
				t.Errorf("takeCharacters() = %v, want = %v", got, test.want)
			}
		})
	}
}
