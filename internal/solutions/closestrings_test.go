package solutions

import "testing"

func TestCloseStrings(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeStrings(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("closeStrings() = %v, want = %v", got, tt.want)
			}
		})
	}
}
