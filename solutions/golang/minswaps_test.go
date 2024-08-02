package golang

import "testing"

func TestMinSwaps(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 0, 1, 1, 0, 0}}, want: 1},
		{name: "smoke 2", args: args{nums: []int{0, 1, 1, 1, 0, 0, 1, 1, 0}}, want: 2},
		{name: "smoke 3", args: args{nums: []int{1, 1, 0, 0, 1}}, want: 0},
		{name: "test 6: runtime", args: args{nums: []int{0, 0, 0}}, want: 0},
		{name: "test 18: wrong answer", args: args{nums: []int{1, 1, 1, 0, 0, 1, 0, 1, 1, 0}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minSwaps(test.args.nums); got != test.want {
				t.Errorf("minSwaps() = %v, want = %v", got, test.want)
			}
		})
	}
}
