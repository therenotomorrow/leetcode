package golang

import "testing"

func TestCountCharacters(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countCharacters(test.args.words, test.args.chars); got != test.want {
				t.Errorf("countCharacters() = %v, want = %v", got, test.want)
			}
		})
	}
}
