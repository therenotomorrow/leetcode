package golang

import "testing"

func TestSumSubarrayMins(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{3, 1, 2, 4}}, want: 17},
		{name: "smoke 2", args: args{arr: []int{11, 81, 94, 43, 3}}, want: 444},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sumSubarrayMins(test.args.arr); got != test.want {
				t.Errorf("sumSubarrayMins() = %v, want = %v", got, test.want)
			}
		})
	}
}
