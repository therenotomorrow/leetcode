package solutions

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "smoke 2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: -1},
		{name: "test 24: wrong answer", args: args{nums: []int{5}, target: 5}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want = %v", got, tt.want)
			}
		})
	}
}
