package golang

import "testing"

func TestNumberOfAlternatingGroups(t *testing.T) {
	t.Parallel()

	type args struct {
		colors []int
		k      int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{colors: []int{0, 1, 0, 1, 0}, k: 3}, want: 3},
		{name: "smoke 2", args: args{colors: []int{0, 1, 0, 0, 1, 0, 1}, k: 6}, want: 2},
		{name: "smoke 3", args: args{colors: []int{1, 1, 0, 1}, k: 4}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfAlternatingGroups(test.args.colors, test.args.k); got != test.want {
				t.Errorf("numberOfAlternatingGroups() = %v, want = %v", got, test.want)
			}
		})
	}
}
