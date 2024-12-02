package golang

import (
	"testing"
)

func TestIsPrefixOfWord(t *testing.T) {
	t.Parallel()

	type args struct {
		sentence   string
		searchWord string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{sentence: "i love eating burger", searchWord: "burg"}, want: 4},
		{name: "smoke 2", args: args{sentence: "this problem is an easy problem", searchWord: "pro"}, want: 2},
		{name: "smoke 3", args: args{sentence: "i am tired", searchWord: "you"}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPrefixOfWord(test.args.sentence, test.args.searchWord); got != test.want {
				t.Errorf("isPrefixOfWord() = %v, want = %v", got, test.want)
			}
		})
	}
}
