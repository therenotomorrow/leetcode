package golang

import "testing"

func TestEliminateMaximum(t *testing.T) {
	t.Parallel()

	type args struct {
		dist  []int
		speed []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{dist: []int{1, 3, 4}, speed: []int{1, 1, 1}}, want: 3},
		{name: "smoke 2", args: args{dist: []int{1, 1, 2, 3}, speed: []int{1, 1, 1, 1}}, want: 1},
		{name: "smoke 3", args: args{dist: []int{3, 2, 4}, speed: []int{5, 3, 2}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := eliminateMaximum(test.args.dist, test.args.speed); got != test.want {
				t.Errorf("eliminateMaximum() = %v, want = %v", got, test.want)
			}
		})
	}
}
