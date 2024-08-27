package golang

import "testing"

func TestCanWinNim(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 4}, want: false},
		{name: "smoke 2", args: args{n: 1}, want: true},
		{name: "smoke 3", args: args{n: 2}, want: true},
		{name: "test 38: wrong answer", args: args{n: 9}, want: true},
		{name: "test 45: wrong answer", args: args{n: 6}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canWinNim(test.args.n); got != test.want {
				t.Errorf("canWinNim() = %v, want = %v", got, test.want)
			}
		})
	}
}
