package golang

import "testing"

func TestMaxProduct(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{3, 4, 5, 2}}, want: 12},
		{name: "smoke 2", args: args{nums: []int{1, 5, 4, 5}}, want: 16},
		{name: "smoke 3", args: args{nums: []int{3, 7}}, want: 12},
		{name: "test 99: wrong answer", args: args{nums: []int{10, 2, 5, 2}}, want: 36},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want = %v", got, tt.want)
			}
		})
	}
}
