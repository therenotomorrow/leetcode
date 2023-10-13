package counting

import "testing"

func Test_findLongestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "eceba", k: 2}, want: 3},
		{args: args{s: "", k: 2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("findLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
