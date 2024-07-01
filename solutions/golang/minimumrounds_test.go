package golang

import "testing"

func TestMinimumRounds(t *testing.T) {
	t.Parallel()

	type args struct {
		tasks []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{tasks: []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}}, want: 4},
		{name: "smoke 2", args: args{tasks: []int{2, 3, 3}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumRounds(test.args.tasks); got != test.want {
				t.Errorf("minimumRounds() = %v, want = %v", got, test.want)
			}
		})
	}
}
