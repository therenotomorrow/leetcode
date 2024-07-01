package golang

import "testing"

func TestAddDigits(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{num: 38}, want: 2},
		{name: "smoke 2", args: args{num: 0}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := addDigits(test.args.num); got != test.want {
				t.Errorf("addDigits() = %v, want = %v", got, test.want)
			}
		})
	}
}
