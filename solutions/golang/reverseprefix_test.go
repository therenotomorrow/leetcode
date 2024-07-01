package golang

import "testing"

func TestReversePrefix(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
		ch   byte
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{word: "abcdefd", ch: 'd'}, want: "dcbaefd"},
		{name: "smoke 2", args: args{word: "xyxzxe", ch: 'z'}, want: "zxyxxe"},
		{name: "smoke 3", args: args{word: "abcd", ch: 'z'}, want: "abcd"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reversePrefix(test.args.word, test.args.ch); got != test.want {
				t.Errorf("reversePrefix() = %v, want = %v", got, test.want)
			}
		})
	}
}
