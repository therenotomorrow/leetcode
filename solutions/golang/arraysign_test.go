package golang

import "testing"

func TestArraySign(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-1, -2, -3, -4, 3, 2, 1}}, want: 1},
		{name: "smoke 2", args: args{nums: []int{1, 5, 0, 2, -3}}, want: 0},
		{name: "smoke 3", args: args{nums: []int{-1, 1, -1, 1, -1}}, want: -1},
		{
			name: "test 31: wrong answer",
			args: args{nums: []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}},
			want: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := arraySign(test.args.nums); got != test.want {
				t.Errorf("arraySign() = %v, want = %v", got, test.want)
			}
		})
	}
}
