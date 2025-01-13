package golang

import "testing"

func TestCanBeValid(t *testing.T) {
	t.Parallel()

	type args struct {
		s      string
		locked string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s: "))()))", locked: "010100"}, want: true},
		{name: "smoke 2", args: args{s: "()()", locked: "0000"}, want: true},
		{name: "smoke 3", args: args{s: ")", locked: "0"}, want: false},
		{
			name: "test 204: wrong answer",
			args: args{s: "((()(()()))()((()()))))()((()(()", locked: "10111100100101001110100010001001"},
			want: true,
		},
		{
			name: "test 207: wrong answer",
			args: args{
				s:      "())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(",
				locked: "100011110110011011010111100111011101111110000101001101001111",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canBeValid(test.args.s, test.args.locked); got != test.want {
				t.Errorf("canBeValid() = %v, want = %v", got, test.want)
			}
		})
	}
}
