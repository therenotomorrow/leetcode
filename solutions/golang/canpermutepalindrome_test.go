package golang

import "testing"

func TestCanPermutePalindrome(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "code"}, want: false},
		{name: "smoke 2", args: args{s: "aab"}, want: true},
		{name: "smoke 3", args: args{s: "carerac"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canPermutePalindrome(test.args.s); got != test.want {
				t.Errorf("canPermutePalindrome() = %v, want = %v", got, test.want)
			}
		})
	}
}
