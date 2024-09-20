package golang

import "testing"

func TestLargestNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{nums: []int{10, 2}}, want: "210"},
		{name: "smoke 2", args: args{nums: []int{3, 30, 34, 5, 9}}, want: "9534330"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := largestNumber(test.args.nums); got != test.want {
				t.Errorf("largestNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
