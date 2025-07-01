package golang

import "testing"

func TestFindLHS(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 2, 2, 5, 2, 3, 7}}, want: 5},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4}}, want: 2},
		{name: "smoke 3", args: args{nums: []int{1, 1, 1, 1}}, want: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findLHS(test.args.nums); got != test.want {
				t.Errorf("findLHS() = %v, want = %v", got, test.want)
			}
		})
	}
}
