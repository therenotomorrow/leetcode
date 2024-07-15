package golang

import "testing"

func TestScoreOfString(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "hello"}, want: 13},
		{name: "smoke 2", args: args{s: "zaz"}, want: 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := scoreOfString(test.args.s); got != test.want {
				t.Errorf("scoreOfString() = %v, want = %v", got, test.want)
			}
		})
	}
}
