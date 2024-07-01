package golang

import "testing"

func TestCloseStrings(t *testing.T) {
	t.Parallel()

	type args struct {
		word1 string
		word2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{word1: "abc", word2: "bca"}, want: true},
		{name: "smoke 2", args: args{word1: "a", word2: "aa"}, want: false},
		{name: "smoke 3", args: args{word1: "cabbba", word2: "abbccc"}, want: true},
		{name: "test 131: wrong answer", args: args{word1: "uau", word2: "ssx"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := closeStrings(test.args.word1, test.args.word2); got != test.want {
				t.Errorf("closeStrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
