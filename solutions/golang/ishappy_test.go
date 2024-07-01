package golang

import "testing"

func TestIsHappy(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{n: 19}, want: true},
		{name: "smoke 2", args: args{n: 2}, want: false},
		{name: "test 10: wrong answer", args: args{n: 7}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isHappy(test.args.n); got != test.want {
				t.Errorf("isHappy() = %v, want = %v", got, test.want)
			}
		})
	}
}
