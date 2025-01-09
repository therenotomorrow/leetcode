package golang

import "testing"

func TestPrefixCount(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
		pref  string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{words: []string{"pay", "attention", "practice", "attend"}, pref: "at"}, want: 2},
		{name: "smoke 2", args: args{words: []string{"leetcode", "win", "loops", "success"}, pref: "code"}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := prefixCount(test.args.words, test.args.pref); got != test.want {
				t.Errorf("prefixCount() = %v, want = %v", got, test.want)
			}
		})
	}
}
