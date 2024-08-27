package golang

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{words: []string{"Hello", "Alaska", "Dad", "Peace"}}, want: []string{"Alaska", "Dad"}},
		{name: "smoke 2", args: args{words: []string{"omk"}}, want: []string{}},
		{name: "smoke 3", args: args{words: []string{"adsdf", "sfd"}}, want: []string{"adsdf", "sfd"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findWords(test.args.words); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findWords() = %v, want = %v", got, test.want)
			}
		})
	}
}
