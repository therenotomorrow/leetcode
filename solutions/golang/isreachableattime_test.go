package golang

import "testing"

func TestIsReachableAtTime(t *testing.T) {
	t.Parallel()

	type args struct {
		sx int
		sy int
		fx int
		fy int
		t  int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{sx: 2, sy: 4, fx: 7, fy: 7, t: 6}, want: true},
		{name: "smoke 2", args: args{sx: 3, sy: 1, fx: 7, fy: 3, t: 3}, want: false},
		{name: "test 769", args: args{sx: 1, sy: 1, fx: 2, fy: 2, t: 1}, want: true},
		{name: "test 799", args: args{sx: 1, sy: 2, fx: 1, fy: 2, t: 1}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isReachableAtTime(test.args.sx, test.args.sy, test.args.fx, test.args.fy, test.args.t); got != test.want {
				t.Errorf("isReachableAtTime() = %v, want = %v", got, test.want)
			}
		})
	}
}
