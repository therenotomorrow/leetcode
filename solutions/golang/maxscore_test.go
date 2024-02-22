package golang

import "testing"

func TestMaxScore(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "011101"}, want: 5},
		{name: "smoke 2", args: args{s: "00111"}, want: 5},
		{name: "smoke 3", args: args{s: "1111"}, want: 3},
		{name: "test 96: wrong answer", args: args{s: "00"}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.s); got != tt.want {
				t.Errorf("maxScore() = %v, want = %v", got, tt.want)
			}
		})
	}
}
