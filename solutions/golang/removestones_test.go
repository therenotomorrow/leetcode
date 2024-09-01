package golang

import "testing"

func TestRemoveStones(t *testing.T) {
	t.Parallel()

	type args struct {
		stones [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}}, want: 5},
		{name: "smoke 2", args: args{stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}}, want: 3},
		{name: "smoke 3", args: args{stones: [][]int{{0, 0}}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := removeStones(test.args.stones); got != test.want {
				t.Errorf("removeStones() = %v, want = %v", got, test.want)
			}
		})
	}
}
