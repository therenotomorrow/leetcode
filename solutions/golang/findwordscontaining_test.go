package golang

import (
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
		x     byte
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{words: []string{"leet", "code"}, x: 'e'}, want: []int{0, 1}},
		{name: "smoke 2", args: args{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'a'}, want: []int{0, 2}},
		{name: "smoke 3", args: args{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'z'}, want: []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findWordsContaining(test.args.words, test.args.x); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findWordsContaining() = %v, want = %v", got, test.want)
			}
		})
	}
}
