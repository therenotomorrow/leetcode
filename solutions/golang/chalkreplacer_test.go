package golang

import "testing"

func TestChalkReplacer(t *testing.T) {
	t.Parallel()

	type args struct {
		chalk []int
		k     int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{chalk: []int{5, 1, 5}, k: 22}, want: 0},
		{name: "smoke 2", args: args{chalk: []int{3, 4, 1, 2}, k: 25}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := chalkReplacer(test.args.chalk, test.args.k); got != test.want {
				t.Errorf("chalkReplacer() = %v, want = %v", got, test.want)
			}
		})
	}
}
