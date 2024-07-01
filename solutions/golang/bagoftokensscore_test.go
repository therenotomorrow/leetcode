package golang

import "testing"

func TestBagOfTokensScore(t *testing.T) {
	t.Parallel()

	type args struct {
		tokens []int
		power  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{tokens: []int{100}, power: 50}, want: 0},
		{name: "smoke 2", args: args{tokens: []int{200, 100}, power: 150}, want: 1},
		{name: "smoke 3", args: args{tokens: []int{100, 200, 300, 400}, power: 200}, want: 2},
		{name: "test 6: wrong answer", args: args{tokens: []int{26}, power: 51}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := bagOfTokensScore(test.args.tokens, test.args.power); got != test.want {
				t.Errorf("bagOfTokensScore() = %v, want = %v", got, test.want)
			}
		})
	}
}
