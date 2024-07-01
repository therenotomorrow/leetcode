package golang

import "testing"

func TestMaxProductDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 6, 2, 7, 4}}, want: 34},
		{name: "smoke 2", args: args{nums: []int{4, 2, 5, 9, 7, 4, 8}}, want: 64},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxProductDifference(test.args.nums); got != test.want {
				t.Errorf("maxProductDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
