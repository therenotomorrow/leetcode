package golang

import "testing"

func TestCountBadPairs(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "smoke 1", args: args{nums: []int{4, 1, 3, 3}}, want: 5},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4, 5}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countBadPairs(test.args.nums); got != test.want {
				t.Errorf("countBadPairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
