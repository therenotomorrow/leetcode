package golang

import "testing"

func TestIsValidPalindrome(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "abcdeca", k: 2}, want: true},
		{name: "smoke 2", args: args{s: "abbababa", k: 1}, want: true},
		{name: "test 3: runtime", args: args{s: "baaaabaa", k: 3}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isValidPalindrome(test.args.s, test.args.k); got != test.want {
				t.Errorf("isValidPalindrome() = %v, want = %v", got, test.want)
			}
		})
	}
}
