package solutions

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{text1: "abcde", text2: "ace"}, want: 3},
		{name: "smoke 2", args: args{text1: "abc", text2: "abc"}, want: 3},
		{name: "smoke 3", args: args{text1: "abc", text2: "def"}, want: 0},
		{name: "test 14: wrong answer", args: args{text1: "hofubmnylkra", text2: "pqhgxgdofcvmr"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want = %v", got, tt.want)
			}
		})
	}
}
