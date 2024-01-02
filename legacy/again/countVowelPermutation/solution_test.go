package countVowelPermutation

import "testing"

func TestCountVowelPermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 1}, want: 5},
		{name: "smoke 2", args: args{n: 2}, want: 10},
		{name: "smoke 3", args: args{n: 5}, want: 68},
		{name: "test 7: wrong answer", args: args{n: 144}, want: 18208803},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelPermutation(tt.args.n); got != tt.want {
				t.Errorf("countVowelPermutation() = %v, want = %v", got, tt.want)
			}
		})
	}
}
