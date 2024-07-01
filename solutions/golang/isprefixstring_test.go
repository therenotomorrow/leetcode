package golang

import "testing"

func TestIsPrefixString(t *testing.T) {
	t.Parallel()

	type args struct {
		s     string
		words []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "iloveleetcode", words: []string{"i", "love", "leetcode", "apples"}}, want: true},
		{name: "smoke 2", args: args{s: "iloveleetcode", words: []string{"apples", "i", "love", "leetcode"}}, want: false},
		{name: "test 216: wrong answer", args: args{s: "a", words: []string{"a", "ad", "cookie"}}, want: true},
		{name: "test 338: wrong answer", args: args{s: "a", words: []string{"aa", "aaaa", "banana"}}, want: false},
		{
			name: "test 339: wrong answer",
			args: args{s: "aaaaaaa", words: []string{"a", "a", "a", "a", "a", "a", "a"}},
			want: true,
		},
		{name: "test 340: wrong answer", args: args{s: "ccccccccc", words: []string{"c", "cc"}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPrefixString(test.args.s, test.args.words); got != test.want {
				t.Errorf("isPrefixString() = %v, want = %v", got, test.want)
			}
		})
	}
}
