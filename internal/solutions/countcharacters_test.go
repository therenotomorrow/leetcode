package solutions

import "testing"

func TestCountCharacters(t *testing.T) {
	type args struct {
		words []string
		chars string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{words: []string{"cat", "bt", "hat", "tree"}, chars: "atach"}, want: 6},
		{name: "smoke 2", args: args{words: []string{"hello", "world", "leetcode"}, chars: "welldonehoneyr"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCharacters(tt.args.words, tt.args.chars); got != tt.want {
				t.Errorf("countCharacters() = %v, want = %v", got, tt.want)
			}
		})
	}
}
