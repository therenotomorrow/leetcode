package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupStrings(t *testing.T) {
	t.Parallel()

	type args struct {
		strings []string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "smoke 1",
			args: args{strings: []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}},
			want: [][]string{{"a", "z"}, {"abc", "bcd", "xyz"}, {"acef"}, {"az", "ba"}},
		},
		{
			name: "smoke 2",
			args: args{strings: []string{"a"}},
			want: [][]string{{"a"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := groupStrings(test.args.strings)
			sort.SliceStable(got, func(i, j int) bool {
				return got[i][0] < got[j][0]
			})

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("groupStrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
