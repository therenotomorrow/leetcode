package golang

import "testing"

func TestIsPalindrome1(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{x: 121}, want: true},
		{name: "smoke 2", args: args{x: -121}, want: false},
		{name: "smoke 3", args: args{x: 10}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPalindrome1(test.args.x); got != test.want {
				t.Errorf("isPalindrome1() = %v, want = %v", got, test.want)
			}
		})
	}
}
