package golang

import "testing"

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "eceba"}, want: 3},
		{name: "smoke 2", args: args{s: "ccaabbb"}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lengthOfLongestSubstringTwoDistinct(test.args.s); got != test.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct() = %v, want = %v", got, test.want)
			}
		})
	}
}
