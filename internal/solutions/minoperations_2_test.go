package solutions

import "testing"

func TestMinOperations2(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations2(tt.args.nums); got != tt.want {
				t.Errorf("minOperations2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
