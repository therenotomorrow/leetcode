package solutions

import "testing"

func TestMaxProductDifference(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{5, 6, 2, 7, 4}}, want: 34},
		{name: "smoke 2", args: args{nums: []int{4, 2, 5, 9, 7, 4, 8}}, want: 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductDifference(tt.args.nums); got != tt.want {
				t.Errorf("maxProductDifference() = %v, want = %v", got, tt.want)
			}
		})
	}
}
