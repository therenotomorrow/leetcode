package golang

import "testing"

func TestShortestDistance(t *testing.T) {
	t.Parallel()

	type args struct {
		wordsDict []string
		word1     string
		word2     string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1:     "coding",
				word2:     "practice",
			},
			want: 3,
		},
		{
			name: "smoke 2",
			args: args{
				wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1:     "makes",
				word2:     "coding",
			},
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := shortestDistance(test.args.wordsDict, test.args.word1, test.args.word2); got != test.want {
				t.Errorf("shortestDistance() = %v, want = %v", got, test.want)
			}
		})
	}
}
