package golang

import "testing"

func TestMySqrt(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{x: 4}, want: 2},
		{name: "smoke 2", args: args{x: 8}, want: 2},
		{name: "test 1015: wrong answer", args: args{x: 0}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := mySqrt(test.args.x); got != test.want {
				t.Errorf("mySqrt() = %v, want = %v", got, test.want)
			}
		})
	}
}
