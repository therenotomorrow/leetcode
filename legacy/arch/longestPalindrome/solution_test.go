package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "babad"}, want: "bab"},
		{name: "smoke 2", args: args{s: "cbbd"}, want: "bb"},
		{name: "test 128: wrong answer", args: args{s: "bb"}, want: "bb"},
		{name: "test 130: wrong answer", args: args{s: "abb"}, want: "bb"},
		{name: "test 133: wrong answer", args: args{s: "caba"}, want: "aba"},
		{name: "test 140: wrong answer", args: args{s: "aaaabaaa"}, want: "aaabaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want = %v", got, tt.want)
			}
		})
	}
}
