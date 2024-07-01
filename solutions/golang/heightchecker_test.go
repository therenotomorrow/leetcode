package golang

import "testing"

func TestHeightChecker(t *testing.T) {
	t.Parallel()

	type args struct {
		heights []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{heights: []int{1, 1, 4, 2, 1, 3}}, want: 3},
		{name: "smoke 2", args: args{heights: []int{5, 1, 2, 3, 4}}, want: 5},
		{name: "smoke 3", args: args{heights: []int{1, 2, 3, 4, 5}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := heightChecker(test.args.heights); got != test.want {
				t.Errorf("heightChecker() = %v, want = %v", got, test.want)
			}
		})
	}
}
