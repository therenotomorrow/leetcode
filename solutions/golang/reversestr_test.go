package golang

import "testing"

func TestReverseStr(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "abcdefg", k: 2}, want: "bacdfeg"},
		{name: "smoke 2", args: args{s: "abcd", k: 2}, want: "bacd"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reverseStr(test.args.s, test.args.k); got != test.want {
				t.Errorf("reverseStr() = %v, want = %v", got, test.want)
			}
		})
	}
}
