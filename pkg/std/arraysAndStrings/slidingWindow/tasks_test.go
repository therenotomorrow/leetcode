package slidingWindow

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, 1, 2, 7, 4, 2, 1, 1, 5}, k: 8}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLengthString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "1101100111"}, want: 5},
		{args: args{s: "11001011"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthString(tt.args.s); got != tt.want {
				t.Errorf("findLengthString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findBestSubarray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, -1, 4, 12, -8, 5, 6}, k: 4}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestSubarray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findBestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
