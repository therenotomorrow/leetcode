package golang

import "testing"

func TestIsStrobogrammatic(t *testing.T) {
	t.Parallel()

	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{num: "69"}, want: true},
		{name: "smoke 2", args: args{num: "88"}, want: true},
		{name: "smoke 3", args: args{num: "962"}, want: false},
		{name: "test 30: wrong answer", args: args{num: "2"}, want: false},
		{name: "test 38: wrong answer", args: args{num: "3"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isStrobogrammatic(test.args.num); got != test.want {
				t.Errorf("isStrobogrammatic() = %v, want = %v", got, test.want)
			}
		})
	}
}
