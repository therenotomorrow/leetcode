package golang

import "testing"

func TestMaxRepeating(t *testing.T) {
	t.Parallel()

	type args struct {
		sequence string
		word     string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{sequence: "ababc", word: "ab"}, want: 2},
		{name: "smoke 2", args: args{sequence: "ababc", word: "ba"}, want: 1},
		{name: "smoke 3", args: args{sequence: "ababc", word: "ac"}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxRepeating(test.args.sequence, test.args.word); got != test.want {
				t.Errorf("maxRepeating() = %v, want = %v", got, test.want)
			}
		})
	}
}
