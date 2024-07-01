package golang

import "testing"

func TestFindMaxConsecutiveOnes1(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 1, 0, 1, 1, 1}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, 0, 1, 1, 0, 1}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMaxConsecutiveOnes1(test.args.nums); got != test.want {
				t.Errorf("findMaxConsecutiveOnes1() = %v, want = %v", got, test.want)
			}
		})
	}
}
