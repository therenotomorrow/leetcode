package golang

import (
	"reflect"
	"testing"
)

func TestGetLongestSubsequence(t *testing.T) {
	t.Parallel()

	type args struct {
		words  []string
		groups []int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{words: []string{"e", "a", "b"}, groups: []int{0, 0, 1}},
			want: []string{"e", "b"},
		},
		{
			name: "smoke 2",
			args: args{words: []string{"a", "b", "c", "d"}, groups: []int{1, 0, 1, 1}},
			want: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getLongestSubsequence(test.args.words, test.args.groups); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getLongestSubsequence() = %v, want = %v", got, test.want)
			}
		})
	}
}
