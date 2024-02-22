package golang

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aabca"}, want: 3},
		{name: "smoke 2", args: args{s: "adc"}, want: 0},
		{name: "smoke 3", args: args{s: "bbcbaba"}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.args.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want = %v", got, tt.want)
			}
		})
	}
}
