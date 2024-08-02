package golang

import "testing"

func TestLongestRepeatingSubstring(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abcd"}, want: 0},
		{name: "smoke 2", args: args{s: "abbaba"}, want: 2},
		{name: "smoke 3", args: args{s: "aabcaabdaab"}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := longestRepeatingSubstring(test.args.s); got != test.want {
				t.Errorf("longestRepeatingSubstring() = %v, want = %v", got, test.want)
			}
		})
	}
}
