package golang

import "testing"

func TestMinDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 3, 2, 4}}, want: 0},
		{name: "smoke 2", args: args{nums: []int{1, 5, 0, 10, 14}}, want: 1},
		{name: "smoke 3", args: args{nums: []int{3, 100, 20}}, want: 0},
		{name: "test 23: wrong answer", args: args{nums: []int{6, 6, 0, 1, 1, 4, 6}}, want: 2},
		{name: "test 44: wrong answer", args: args{nums: []int{82, 81, 95, 75, 20}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minDifference(test.args.nums); got != test.want {
				t.Errorf("minDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
