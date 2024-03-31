package golang

import "testing"

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringKDistinct(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("lengthOfLongestSubstringKDistinct() = %v, want = %v", got, tt.want)
			}
		})
	}
}
