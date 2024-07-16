package golang

import "testing"

func TestMaximumNumberOfStringPairs(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{words: []string{"cd", "ac", "dc", "ca", "zz"}}, want: 2},
		{name: "smoke 2", args: args{words: []string{"ab", "ba", "cc"}}, want: 1},
		{name: "smoke 3", args: args{words: []string{"aa", "ab"}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumNumberOfStringPairs(test.args.words); got != test.want {
				t.Errorf("maximumNumberOfStringPairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
