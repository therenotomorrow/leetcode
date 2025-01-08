package golang

import "testing"

func TestCountPrefixSuffixPairs(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{words: []string{"a", "aba", "ababa", "aa"}}, want: 4},
		{name: "smoke 2", args: args{words: []string{"pa", "papa", "ma", "mama"}}, want: 2},
		{name: "smoke 3", args: args{words: []string{"abab", "ab"}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countPrefixSuffixPairs(test.args.words); got != test.want {
				t.Errorf("countPrefixSuffixPairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
