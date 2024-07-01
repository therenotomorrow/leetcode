package golang

import "testing"

func TestToHex(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{num: 26}, want: "1a"},
		{name: "smoke 2", args: args{num: -1}, want: "ffffffff"},
		{name: "test 99: wrong answer", args: args{num: 0}, want: "0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := toHex(test.args.num); got != test.want {
				t.Errorf("toHex() = %v, want = %v", got, test.want)
			}
		})
	}
}
