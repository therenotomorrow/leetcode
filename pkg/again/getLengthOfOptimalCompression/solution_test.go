package getLengthOfOptimalCompression

import "testing"

func TestGetLengthOfOptimalCompression(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "aaabcccd", k: 2}, want: 4},
		{name: "smoke 2", args: args{s: "aabbaa", k: 2}, want: 2},
		{name: "smoke 3", args: args{s: "aaaaaaaaaaa", k: 0}, want: 3},
		{name: "test 83: tle", args: args{s: "bcaaaacbabdabdadbababadcdcbcdababbbcbdadacdaaddccadbcabcbbacaaaddddcacdccdbdbab", k: 58}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLengthOfOptimalCompression(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getLengthOfOptimalCompression() = %v, want = %v", got, tt.want)
			}
		})
	}
}
