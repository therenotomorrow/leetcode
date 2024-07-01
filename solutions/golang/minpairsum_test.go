package golang

import "testing"

func TestMinPairSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{3, 5, 2, 3}}, want: 7},
		{name: "smoke 2", args: args{nums: []int{3, 5, 4, 2, 4, 6}}, want: 8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minPairSum(test.args.nums); got != test.want {
				t.Errorf("minPairSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
