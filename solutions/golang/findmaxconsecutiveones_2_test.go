package golang

import "testing"

func TestFindMaxConsecutiveOnes2(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 0, 1, 1, 0}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{1, 0, 1, 1, 0, 1}}, want: 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMaxConsecutiveOnes2(test.args.nums); got != test.want {
				t.Errorf("findMaxConsecutiveOnes2() = %v, want = %v", got, test.want)
			}
		})
	}
}
