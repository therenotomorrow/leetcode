package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestStringMatching(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{words: []string{"mass", "as", "hero", "superhero"}}, want: []string{"as", "hero"}},
		{name: "smoke 2", args: args{words: []string{"leetcode", "et", "code"}}, want: []string{"code", "et"}},
		{name: "smoke 3", args: args{words: []string{"blue", "green", "bu"}}, want: []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := stringMatching(test.args.words)

			sort.Strings(got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("stringMatching() = %v, want = %v", got, test.want)
			}
		})
	}
}
