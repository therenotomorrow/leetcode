package solutions

import "testing"

func TestKthGrammar(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 1, k: 1}, want: 0},
		{name: "smoke 2", args: args{n: 2, k: 1}, want: 0},
		{name: "smoke 3", args: args{n: 2, k: 2}, want: 1},
		{name: "test 15: oom", args: args{n: 30, k: 434991989}, want: 0},
		{name: "test 32: wrong answer", args: args{n: 3, k: 3}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthGrammar() = %v, want = %v", got, tt.want)
			}
		})
	}
}
