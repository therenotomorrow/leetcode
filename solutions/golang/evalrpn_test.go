package golang

import "testing"

func TestEvalRPN(t *testing.T) {
	t.Parallel()

	type args struct {
		tokens []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{tokens: []string{"2", "1", "+", "3", "*"}}, want: 9},
		{name: "smoke 2", args: args{tokens: []string{"4", "13", "5", "/", "+"}}, want: 6},
		{
			name: "smoke 3",
			args: args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			want: 22,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := evalRPN(test.args.tokens); got != test.want {
				t.Errorf("evalRPN() = %v, want = %v", got, test.want)
			}
		})
	}
}
