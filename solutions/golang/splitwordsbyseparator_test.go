package golang

import (
	"reflect"
	"testing"
)

func TestSplitWordsBySeparator(t *testing.T) {
	t.Parallel()

	type args struct {
		words     []string
		separator byte
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{words: []string{"one.two.three", "four.five", "six"}, separator: '.'},
			want: []string{"one", "two", "three", "four", "five", "six"},
		},
		{
			name: "smoke 2",
			args: args{words: []string{"$easy$", "$problem$"}, separator: '$'},
			want: []string{"easy", "problem"},
		},
		{
			name: "smoke 3",
			args: args{words: []string{"|||"}, separator: '|'},
			want: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := splitWordsBySeparator(test.args.words, test.args.separator); !reflect.DeepEqual(got, test.want) {
				t.Errorf("splitWordsBySeparator() = %v, want = %v", got, test.want)
			}
		})
	}
}
