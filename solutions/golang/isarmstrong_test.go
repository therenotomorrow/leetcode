package golang

import "testing"

func TestIsArmstrong(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 153}, want: true},
		{name: "smoke 2", args: args{n: 123}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isArmstrong(test.args.n); got != test.want {
				t.Errorf("isArmstrong() = %v, want = %v", got, test.want)
			}
		})
	}
}
