package golang

import "testing"

func TestMinimumChairs(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{"EEEEEEE"}, want: 7},
		{name: "smoke 2", args: args{"ELELEEL"}, want: 2},
		{name: "smoke 3", args: args{"ELEELEELLL"}, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumChairs(test.args.s); got != test.want {
				t.Errorf("minimumChairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
