package golang

import "testing"

func TestCalculateTime(t *testing.T) {
	t.Parallel()

	type args struct {
		keyboard string
		word     string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{keyboard: "abcdefghijklmnopqrstuvwxyz", word: "cba"}, want: 4},
		{name: "smoke 2", args: args{keyboard: "pqrstuvwxyzabcdefghijklmno", word: "leetcode"}, want: 73},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := calculateTime(test.args.keyboard, test.args.word); got != test.want {
				t.Errorf("calculateTime() = %v, want = %v", got, test.want)
			}
		})
	}
}
