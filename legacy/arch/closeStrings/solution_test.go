package closeStrings

import "testing"

func Test_closeStrings(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{word1: "abc", word2: "bca"}, want: true},
		{args: args{word1: "a", word2: "aa"}, want: false},
		{args: args{word1: "cabbba", word2: "abbccc"}, want: true},
		{args: args{word1: "uau", word2: "ssx"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeStrings(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("closeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
