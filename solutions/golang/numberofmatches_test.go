package golang

import "testing"

func TestNumberOfMatches(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 7}, want: 6},
		{name: "smoke 2", args: args{n: 14}, want: 13},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfMatches(test.args.n); got != test.want {
				t.Errorf("numberOfMatches() = %v, want = %v", got, test.want)
			}
		})
	}
}
