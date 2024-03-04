package golang

import "testing"

func TestBagOfTokensScore(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagOfTokensScore(tt.args.tokens, tt.args.power); got != tt.want {
				t.Errorf("bagOfTokensScore() = %v, want = %v", got, tt.want)
			}
		})
	}
}
