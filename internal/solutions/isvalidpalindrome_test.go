package solutions

import "testing"

func TestIsValidPalindrome(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPalindrome(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("isValidPalindrome() = %v, want = %v", got, tt.want)
			}
		})
	}
}
