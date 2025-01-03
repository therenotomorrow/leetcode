package golang

import "testing"

func TestWaysToSplitArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{10, 4, -8, 7}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{2, 3, 1, 0}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := waysToSplitArray(test.args.nums); got != test.want {
				t.Errorf("waysToSplitArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
