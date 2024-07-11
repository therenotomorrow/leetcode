package golang

import "testing"

func TestIsUgly(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 6}, want: true},
		{name: "smoke 2", args: args{n: 1}, want: true},
		{name: "smoke 3", args: args{n: 14}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isUgly(test.args.n); got != test.want {
				t.Errorf("isUgly() = %v, want = %v", got, test.want)
			}
		})
	}
}
