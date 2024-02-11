package solutions

import "testing"

func TestCountSubstrings(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "abc"}, want: 3},
		{name: "smoke 2", args: args{s: "aaa"}, want: 6},
		{name: "test 6: wrong answer", args: args{s: "a"}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want = %v", got, tt.want)
			}
		})
	}
}
