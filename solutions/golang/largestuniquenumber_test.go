package golang

import "testing"

func TestLargestUniqueNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 7, 3, 9, 4, 9, 8, 3, 1}}, want: 8},
		{name: "smoke 2", args: args{nums: []int{9, 9, 8, 8}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestUniqueNumber(test.args.nums); got != test.want {
				t.Errorf("largestUniqueNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
