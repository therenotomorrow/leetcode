package golang

import "testing"

func TestMaxWidthRamp(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{6, 0, 8, 2, 1, 5}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}}, want: 7},
		{name: "test 40: wrong answer", args: args{nums: []int{0, 1}}, want: 1},
		{name: "test 88: wrong answer", args: args{nums: []int{2, 2, 1}}, want: 1},
		{name: "test 93: wrong answer", args: args{nums: []int{2, 3, 1}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maxWidthRamp(test.args.nums); got != test.want {
				t.Errorf("maxWidthRamp() = %v, want = %v", got, test.want)
			}
		})
	}
}
