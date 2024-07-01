package golang

import "testing"

func TestValidPalindrome(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "aba"}, want: true},
		{name: "smoke 2", args: args{s: "abca"}, want: true},
		{name: "smoke 3", args: args{s: "abc"}, want: false},
		{name: "test 427: wrong answer", args: args{s: "eeccccbebaeeabebccceea"}, want: false},
		{name: "test 466: wrong answer", args: args{s: "ebcbbececabbacecbbcbe"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := validPalindrome(test.args.s); got != test.want {
				t.Errorf("validPalindrome() = %v, want = %v", got, test.want)
			}
		})
	}
}
