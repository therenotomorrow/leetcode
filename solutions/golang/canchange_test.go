package golang

import (
	"testing"
)

func TestCanChange(t *testing.T) {
	t.Parallel()

	type args struct {
		start  string
		target string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{start: "_L__R__R_", target: "L______RR"}, want: true},
		{name: "smoke 2", args: args{start: "R_L_", target: "__LR"}, want: false},
		{name: "smoke 3", args: args{start: "_R", target: "R_"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canChange(test.args.start, test.args.target); got != test.want {
				t.Errorf("canChange() = %v, want = %v", got, test.want)
			}
		})
	}
}
