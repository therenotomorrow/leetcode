package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestUncommonFromSentences(t *testing.T) {
	t.Parallel()

	type args struct {
		s1 string
		s2 string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{s1: "this apple is sweet", s2: "this apple is sour"},
			want: []string{"sour", "sweet"},
		},
		{
			name: "smoke 2",
			args: args{s1: "apple apple", s2: "banana"}, //nolint:dupword
			want: []string{"banana"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := uncommonFromSentences(test.args.s1, test.args.s2)

			sort.Strings(got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("uncommonFromSentences() = %v, want = %v", got, test.want)
			}
		})
	}
}
