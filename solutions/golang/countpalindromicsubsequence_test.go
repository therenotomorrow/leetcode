package golang

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countPalindromicSubsequence(test.args.s); got != test.want {
				t.Errorf("countPalindromicSubsequence() = %v, want = %v", got, test.want)
			}
		})
	}
}
