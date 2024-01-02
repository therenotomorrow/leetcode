package solutions

import "testing"

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringTwoDistinct(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct() = %v, want = %v", got, tt.want)
			}
		})
	}
}
