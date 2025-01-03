package golang

import (
	"reflect"
	"testing"
)

func TestVowelStrings(t *testing.T) {
	t.Parallel()

	type args struct {
		words   []string
		queries [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{words: []string{"aba", "bcb", "ece", "aa", "e"}, queries: [][]int{{0, 2}, {1, 4}, {1, 1}}},
			want: []int{2, 3, 0},
		},
		{
			name: "smoke 2",
			args: args{words: []string{"a", "e", "i"}, queries: [][]int{{0, 2}, {0, 1}, {2, 2}}},
			want: []int{3, 2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := vowelStrings(test.args.words, test.args.queries); !reflect.DeepEqual(got, test.want) {
				t.Errorf("vowelStrings() = %v, want = %v", got, test.want)
			}
		})
	}
}
