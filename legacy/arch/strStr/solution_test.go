package strStr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{haystack: "sadbutsad", needle: "sad"}, want: 0},
		{args: args{haystack: "leetcode", needle: "leeto"}, want: -1},
		{args: args{haystack: "a", needle: "a"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
