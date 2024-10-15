package golang

import "testing"

func TestFindTheLongestSubstring(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "eleetminicoworoep"}, want: 13},
		{name: "smoke 2", args: args{s: "leetcodeisgreat"}, want: 5},
		{name: "smoke 3", args: args{s: "bcbcbc"}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findTheLongestSubstring(test.args.s); got != test.want {
				t.Errorf("findTheLongestSubstring() = %v, want = %v", got, test.want)
			}
		})
	}
}
