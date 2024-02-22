package golang

import "testing"

func TestShortestDistance(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistance(tt.args.wordsDict, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("shortestDistance() = %v, want = %v", got, tt.want)
			}
		})
	}
}
