package golang

import "testing"

func TestMaxAscendingSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{10, 20, 30, 5, 10, 50}}, want: 65},
		{name: "smoke 2", args: args{nums: []int{10, 20, 30, 40, 50}}, want: 150},
		{name: "smoke 3", args: args{nums: []int{12, 17, 15, 13, 10, 11, 12}}, want: 33},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxAscendingSum(test.args.nums); got != test.want {
				t.Errorf("maxAscendingSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
