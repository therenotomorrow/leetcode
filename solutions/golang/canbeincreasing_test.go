package golang

import "testing"

func TestCanBeIncreasing(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 10, 5, 7}}, want: true},
		{name: "smoke 2", args: args{nums: []int{2, 3, 1, 2}}, want: false},
		{name: "smoke 3", args: args{nums: []int{1, 1, 1}}, want: false},
		{name: "test 61: wrong answer", args: args{nums: []int{1, 1}}, want: true},
		{name: "test 74: wrong answer", args: args{nums: []int{541, 783, 433, 744}}, want: false},
		{name: "test 77: wrong answer", args: args{nums: []int{89, 384, 691, 680, 111, 756}}, want: false},
		{name: "test 78: wrong answer", args: args{nums: []int{512, 867, 904, 997, 403}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canBeIncreasing(test.args.nums); got != test.want {
				t.Errorf("canBeIncreasing() = %v, want = %v", got, test.want)
			}
		})
	}
}
