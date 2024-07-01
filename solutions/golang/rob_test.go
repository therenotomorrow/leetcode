package golang

import "testing"

func TestRob(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 1}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{2, 7, 9, 3, 1}}, want: 12},
		{name: "own 1", args: args{nums: []int{1, 6, 1, 4, 5}}, want: 11},
		{name: "own 2", args: args{nums: []int{1, 4, 2}}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := rob(test.args.nums); got != test.want {
				t.Errorf("rob() = %v, want = %v", got, test.want)
			}
		})
	}
}
