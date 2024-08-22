package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateAbbreviations(t *testing.T) {
	t.Parallel()

	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{word: "word"},
			want: []string{
				"4", "3d", "2r1", "2rd", "1o2", "1o1d", "1or1", "1ord", "w3", "w2d", "w1r1", "w1rd", "wo2", "wo1d", "wor1", "word",
			},
		},
		{
			name: "smoke 2",
			args: args{word: "a"},
			want: []string{"1", "a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := generateAbbreviations(test.args.word)

			sort.Strings(got)
			sort.Strings(test.want)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("generateAbbreviations() = %v, want = %v", got, test.want)
			}
		})
	}
}
