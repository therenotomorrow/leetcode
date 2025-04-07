package golang

import "testing"

func TestCanPartition(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{1, 5, 11, 5}}, want: true},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 5}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canPartition(test.args.nums); got != test.want {
				t.Errorf("canPartition() = %v, want = %v", got, test.want)
			}
		})
	}
}
