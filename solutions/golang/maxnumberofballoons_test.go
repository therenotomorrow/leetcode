package golang

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	t.Parallel()

	type args struct {
		text string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{text: "nlaebolko"}, want: 1},
		{name: "smoke 2", args: args{text: "loonbalxballpoon"}, want: 2},
		{name: "smoke 3", args: args{text: "leetcode"}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxNumberOfBalloons(test.args.text); got != test.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, test.want)
			}
		})
	}
}
