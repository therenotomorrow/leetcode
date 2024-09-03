package golang

import "testing"

func TestBuddyStrings(t *testing.T) {
	t.Parallel()

	type args struct {
		s    string
		goal string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "ab", goal: "ba"}, want: true},
		{name: "smoke 2", args: args{s: "ab", goal: "ab"}, want: false},
		{name: "smoke 3", args: args{s: "aa", goal: "aa"}, want: true},
		{name: "test 33: wrong answer", args: args{s: "abcaa", goal: "abcbb"}, want: false},
		{name: "test 35: wrong answer", args: args{s: "abcd", goal: "badc"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := buddyStrings(test.args.s, test.args.goal); got != test.want {
				t.Errorf("buddyStrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
