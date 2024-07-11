package golang

import "testing"

func TestThirdMax(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{3, 2, 1}}, want: 1},
		{name: "smoke 2", args: args{nums: []int{1, 2}}, want: 2},
		{name: "smoke 3", args: args{nums: []int{2, 2, 3, 1}}, want: 1},
		{name: "test 22: wrong answer", args: args{nums: []int{1, 1, 2}}, want: 2},
		{name: "test 32: wrong answer", args: args{nums: []int{1, 2, 2, 5, 3, 5}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := thirdMax(test.args.nums); got != test.want {
				t.Errorf("thirdMax() = %v, want = %v", got, test.want)
			}
		})
	}
}
