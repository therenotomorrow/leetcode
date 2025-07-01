package golang

import "testing"

func TestMaxDistance1(t *testing.T) {
	t.Parallel()

	type args struct {
		arrays [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}}, want: 4},
		{name: "smoke 2", args: args{arrays: [][]int{{1}, {1}}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxDistance1(test.args.arrays); got != test.want {
				t.Errorf("maxDistance1() = %v, want = %v", got, test.want)
			}
		})
	}
}
