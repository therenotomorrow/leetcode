package golang

import "testing"

func TestKnightDialer(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 1}, want: 10},
		{name: "smoke 2", args: args{n: 2}, want: 20},
		{name: "smoke 3", args: args{n: 3131}, want: 136006598},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := knightDialer(test.args.n); got != test.want {
				t.Errorf("knightDialer() = %v, want = %v", got, test.want)
			}
		})
	}
}
