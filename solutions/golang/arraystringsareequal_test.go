package golang

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		word1 []string
		word2 []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{word1: []string{"ab", "c"}, word2: []string{"a", "bc"}}, want: true},
		{name: "smoke 2", args: args{word1: []string{"a", "cb"}, word2: []string{"ab", "c"}}, want: false},
		{name: "smoke 3", args: args{word1: []string{"abc", "d", "defg"}, word2: []string{"abcddefg"}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := arrayStringsAreEqual(test.args.word1, test.args.word2); got != test.want {
				t.Errorf("arrayStringsAreEqual() = %v, want = %v", got, test.want)
			}
		})
	}
}
