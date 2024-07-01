package golang

import "testing"

func TestFindMaxK(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-1, 2, -3, 3}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{-1, 10, 6, 7, -7, 1}}, want: 7},
		{name: "smoke 3", args: args{nums: []int{-10, 8, 6, 7, -2, -3}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMaxK(test.args.nums); got != test.want {
				t.Errorf("findMaxK() = %v, want = %v", got, test.want)
			}
		})
	}
}
