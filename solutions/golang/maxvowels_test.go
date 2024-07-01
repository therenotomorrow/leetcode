package golang

import "testing"

func TestMaxVowels(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abciiidef", k: 3}, want: 3},
		{name: "smoke 2", args: args{s: "aeiou", k: 2}, want: 2},
		{name: "smoke 3", args: args{s: "leetcode", k: 3}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxVowels(test.args.s, test.args.k); got != test.want {
				t.Errorf("maxVowels() = %v, want = %v", got, test.want)
			}
		})
	}
}
