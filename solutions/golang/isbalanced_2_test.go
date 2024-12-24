package golang

import (
	"testing"
)

func TestIsBalanced2(t *testing.T) {
	t.Parallel()

	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{num: "1234"}, want: false},
		{name: "smoke 2", args: args{num: "24123"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isBalanced2(test.args.num); got != test.want {
				t.Errorf("isBalanced2() = %v, want = %v", got, test.want)
			}
		})
	}
}
