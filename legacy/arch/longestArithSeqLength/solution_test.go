package longestArithSeqLength

import "testing"

func Test_longestArithSeqLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, 6, 9, 12}}, want: 4},
		{args: args{nums: []int{9, 4, 7, 2, 10}}, want: 3},
		{args: args{nums: []int{20, 1, 15, 3, 10, 5, 8}}, want: 4},
		{args: args{nums: []int{83, 20, 17, 43, 52, 78, 68, 45}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestArithSeqLength(tt.args.nums); got != tt.want {
				t.Errorf("longestArithSeqLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
