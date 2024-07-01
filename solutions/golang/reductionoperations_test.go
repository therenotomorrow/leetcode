package golang

import "testing"

func TestReductionOperations(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 1, 3}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1}}, want: 0},
		{name: "smoke 3", args: args{nums: []int{1, 1, 2, 2, 3}}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reductionOperations(test.args.nums); got != test.want {
				t.Errorf("reductionOperations() = %v, want = %v", got, test.want)
			}
		})
	}
}
