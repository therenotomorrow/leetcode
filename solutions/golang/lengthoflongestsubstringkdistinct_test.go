package golang

import "testing"

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
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
		{name: "smoke 1", args: args{s: "eceba", k: 2}, want: 3},
		{name: "smoke 2", args: args{s: "aa", k: 1}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lengthOfLongestSubstringKDistinct(test.args.s, test.args.k); got != test.want {
				t.Errorf("lengthOfLongestSubstringKDistinct() = %v, want = %v", got, test.want)
			}
		})
	}
}
