package golang

import "testing"

func TestMaxScore(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxScore(test.args.s); got != test.want {
				t.Errorf("maxScore() = %v, want = %v", got, test.want)
			}
		})
	}
}
