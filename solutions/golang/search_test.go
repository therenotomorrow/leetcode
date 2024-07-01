package golang

import "testing"

func TestSearch(t *testing.T) {
	t.Parallel()

	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "smoke 2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: -1},
		{name: "test 24: wrong answer", args: args{nums: []int{5}, target: 5}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := search(test.args.nums, test.args.target); got != test.want {
				t.Errorf("search() = %v, want = %v", got, test.want)
			}
		})
	}
}
