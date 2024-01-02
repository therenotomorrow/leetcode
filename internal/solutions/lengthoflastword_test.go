package solutions

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "Hello World"}, want: 5},
		{name: "smoke 2", args: args{s: "   fly me   to   the moon  "}, want: 4},
		{name: "smoke 3", args: args{s: "luffy is still joyboy"}, want: 6},
		{name: "test 49: wrong answer", args: args{s: "a"}, want: 1},
		{name: "test 51: wrong answer", args: args{s: "day"}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want = %v", got, tt.want)
			}
		})
	}
}
