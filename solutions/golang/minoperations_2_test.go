package golang

import "testing"

func TestMinOperations2(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 3, 3, 2, 2, 4, 2, 3, 4}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{2, 1, 2, 2, 3, 3}}, want: -1},
		{
			name: "test 371: wrong answer",
			args: args{nums: []int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}},
			want: 7,
		},
		{
			name: "test 626: wrong answer",
			args: args{nums: []int{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19}},
			want: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations2(test.args.nums); got != test.want {
				t.Errorf("minOperations2() = %v, want = %v", got, test.want)
			}
		})
	}
}
