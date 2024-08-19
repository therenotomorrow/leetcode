package golang

import "testing"

func TestMinSteps1(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minSteps1(test.args.s, test.args.t); got != test.want {
				t.Errorf("minSteps1() = %v, want = %v", got, test.want)
			}
		})
	}
}
