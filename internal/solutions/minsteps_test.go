package solutions

import "testing"

func TestMinSteps(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "bab", t: "aba"}, want: 1},
		{name: "smoke 2", args: args{s: "leetcode", t: "practice"}, want: 5},
		{name: "smoke 3", args: args{s: "anagram", t: "mangaar"}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minSteps() = %v, want = %v", got, tt.want)
			}
		})
	}
}
