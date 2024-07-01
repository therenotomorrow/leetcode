package golang

import "testing"

func TestFindMiddleIndex(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 3, -1, 8, 4}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{1, -1, 4}}, want: 2},
		{name: "smoke 3", args: args{nums: []int{2, 5}}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMiddleIndex(test.args.nums); got != test.want {
				t.Errorf("findMiddleIndex() = %v, want = %v", got, test.want)
			}
		})
	}
}
