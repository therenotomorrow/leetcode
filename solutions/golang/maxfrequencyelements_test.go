package golang

import "testing"

func TestMaxFrequencyElements(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 2, 3, 1, 4}}, want: 4},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4, 5}}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElements(tt.args.nums); got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want = %v", got, tt.want)
			}
		})
	}
}
