package solutions

import "testing"

func TestCanPermutePalindrome(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome(tt.args.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want = %v", got, tt.want)
			}
		})
	}
}
