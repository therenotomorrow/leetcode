package golang

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	t.Parallel()

	type args struct {
		allowed string
		words   []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{allowed: "ab", words: []string{"ad", "bd", "aaab", "baa", "badab"}},
			want: 2,
		},
		{
			name: "smoke 2",
			args: args{allowed: "abc", words: []string{"a", "b", "c", "ab", "ac", "bc", "abc"}},
			want: 7,
		},
		{
			name: "smoke 3",
			args: args{allowed: "cad", words: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countConsistentStrings(test.args.allowed, test.args.words); got != test.want {
				t.Errorf("countConsistentStrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
