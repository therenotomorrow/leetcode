package golang

import "testing"

func TestLargestPerimeter(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{nums: []int{5, 5, 5}}, want: 15},
		{name: "smoke 2", args: args{nums: []int{1, 12, 1, 2, 5, 50, 3}}, want: 12},
		{name: "smoke 3", args: args{nums: []int{5, 5, 50}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestPerimeter(test.args.nums); got != test.want {
				t.Errorf("largestPerimeter() = %v, want = %v", got, test.want)
			}
		})
	}
}
