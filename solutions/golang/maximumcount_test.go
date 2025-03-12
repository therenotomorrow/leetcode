package golang

import "testing"

func TestMaximumCount(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-2, -1, -1, 1, 2, 3}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2}}, want: 3},
		{name: "smoke 3", args: args{nums: []int{5, 20, 66, 1314}}, want: 4},
		{
			name: "test 89: wrong answer",
			args: args{nums: []int{-1563, -236, -114, -55, 427, 447, 687, 752, 1021, 1636}},
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := maximumCount(test.args.nums); got != test.want {
				t.Errorf("maximumCount() = %v, want = %v", got, test.want)
			}
		})
	}
}
