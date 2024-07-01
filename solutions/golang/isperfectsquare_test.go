package golang

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{num: 16}, want: true},
		{name: "smoke 2", args: args{num: 14}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isPerfectSquare(test.args.num); got != test.want {
				t.Errorf("isPerfectSquare() = %v, want = %v", got, test.want)
			}
		})
	}
}
