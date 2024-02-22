package golang

import "testing"

func TestCalculateTime(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTime(tt.args.keyboard, tt.args.word); got != tt.want {
				t.Errorf("calculateTime() = %v, want = %v", got, tt.want)
			}
		})
	}
}
