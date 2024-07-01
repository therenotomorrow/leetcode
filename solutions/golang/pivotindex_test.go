package golang

import "testing"

func TestPivotIndex(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{[]int{1, 7, 3, 6, 5, 6}}, want: 3},
		{name: "smoke 2", args: args{[]int{1, 2, 3}}, want: -1},
		{name: "smoke 3", args: args{[]int{2, 1, -1}}, want: 0},
		{name: "test 686: wrong answer", args: args{[]int{-1, -1, 0, 1, 1, 0}}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := pivotIndex(test.args.nums); got != test.want {
				t.Errorf("pivotIndex() = %v, want = %v", got, test.want)
			}
		})
	}
}
