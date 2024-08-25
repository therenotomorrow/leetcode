package golang

import "testing"

func TestNearestPalindromic(t *testing.T) {
	t.Parallel()

	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{n: "123"}, want: "121"},
		{name: "smoke 2", args: args{n: "1"}, want: "0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := nearestPalindromic(test.args.n); got != test.want {
				t.Errorf("nearestPalindromic() = %v, want = %v", got, test.want)
			}
		})
	}
}
