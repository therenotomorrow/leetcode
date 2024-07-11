package golang

import "testing"

func TestReverseParentheses(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "(abcd)"}, want: "dcba"},
		{name: "smoke 2", args: args{s: "(u(love)i)"}, want: "iloveu"},
		{name: "smoke 3", args: args{s: "(ed(et(oc))el)"}, want: "leetcode"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reverseParentheses(test.args.s); got != test.want {
				t.Errorf("reverseParentheses() = %v, want = %v", got, test.want)
			}
		})
	}
}
