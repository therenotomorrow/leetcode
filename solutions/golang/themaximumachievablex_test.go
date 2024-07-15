package golang

import "testing"

func TestTheMaximumAchievableX(t *testing.T) {
	t.Parallel()

	type args struct {
		num int
		t   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{num: 4, t: 1}, want: 6},
		{name: "smoke 2", args: args{num: 3, t: 2}, want: 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := theMaximumAchievableX(test.args.num, test.args.t); got != test.want {
				t.Errorf("theMaximumAchievableX() = %v, want = %v", got, test.want)
			}
		})
	}
}
