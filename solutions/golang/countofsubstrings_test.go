package golang

import "testing"

func TestCountOfSubstrings(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
		k    int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{word: "aeioqq", k: 1}, want: 0},
		{name: "smoke 2", args: args{word: "aeiou", k: 0}, want: 1},
		{name: "smoke 3", args: args{word: "ieaouqqieaouqq", k: 1}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countOfSubstrings(test.args.word, test.args.k); got != test.want {
				t.Errorf("countOfSubstrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
