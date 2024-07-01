package golang

import "testing"

func TestKthGrammar(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := kthGrammar(test.args.n, test.args.k); got != test.want {
				t.Errorf("kthGrammar() = %v, want = %v", got, test.want)
			}
		})
	}
}
