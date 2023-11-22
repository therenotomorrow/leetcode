package solutions

import "testing"

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{2, 2, 1, 2, 1}}, want: 2},
		{name: "smoke 2", args: args{arr: []int{100, 1, 1000}}, want: 3},
		{name: "smoke 3", args: args{arr: []int{1, 2, 3, 4, 5}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.args.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want = %v", got, tt.want)
			}
		})
	}
}
