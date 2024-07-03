package golang

import "testing"

func TestIsAlienSorted(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
		order string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "smoke 1",
			args: args{words: []string{"hello", "leetcode"}, order: "hlabcdefgijkmnopqrstuvwxyz"},
			want: true,
		},
		{
			name: "smoke 2",
			args: args{words: []string{"word", "world", "row"}, order: "worldabcefghijkmnpqstuvxyz"},
			want: false,
		},
		{
			name: "smoke 3",
			args: args{words: []string{"apple", "app"}, order: "abcdefghijklmnopqrstuvwxyz"},
			want: false,
		},
		{
			name: "test 111",
			args: args{words: []string{"kuvp", "q"}, order: "ngxlkthsjuoqcpavbfdermiywz"},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isAlienSorted(test.args.words, test.args.order); got != test.want {
				t.Errorf("isAlienSorted() = %v, want = %v", got, test.want)
			}
		})
	}
}
