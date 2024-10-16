package golang

import "testing"

func TestLongestDiverseString(t *testing.T) {
	t.Parallel()

	type args struct {
		a int
		b int
		c int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{a: 1, b: 1, c: 7}, want: "ccaccbcc"},
		{name: "smoke 2", args: args{a: 7, b: 1, c: 0}, want: "aabaa"},
		{name: "test 25: wrong answer", args: args{a: 0, b: 8, c: 11}, want: "ccbccbcbcbcbcbcbcbc"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestDiverseString(test.args.a, test.args.b, test.args.c); got != test.want {
				t.Errorf("longestDiverseString() = %v, want = %v", got, test.want)
			}
		})
	}
}
