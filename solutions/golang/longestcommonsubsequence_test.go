package golang

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestCommonSubsequence(test.args.text1, test.args.text2); got != test.want {
				t.Errorf("longestCommonSubsequence() = %v, want = %v", got, test.want)
			}
		})
	}
}
