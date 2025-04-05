package golang

import "testing"

func TestSubsetXORSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3}}, want: 6},
		{name: "smoke 2", args: args{nums: []int{5, 1, 6}}, want: 28},
		{name: "smoke 3", args: args{nums: []int{3, 4, 5, 6, 7, 8}}, want: 480},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := subsetXORSum(test.args.nums); got != test.want {
				t.Errorf("subsetXORSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
