package golang

import "testing"

func TestMinWindow(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
		{name: "smoke 2", args: args{s: "a", t: "a"}, want: "a"},
		{name: "smoke 3", args: args{s: "a", t: "aa"}, want: ""},
		{name: "test 153: wrong answer", args: args{s: "ab", t: "b"}, want: "b"},
		{name: "test 172: wrong answer", args: args{s: "abc", t: "b"}, want: "b"},
		{name: "test 207: wrong answer", args: args{s: "cabwefgewcwaefgcf", t: "cae"}, want: "cwae"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minWindow(test.args.s, test.args.t); got != test.want {
				t.Errorf("minWindow() = %v, want = %v", got, test.want)
			}
		})
	}
}
