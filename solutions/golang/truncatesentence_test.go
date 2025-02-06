package golang

import "testing"

func TestTruncateSentence(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "Hello how are you Contestant", k: 4}, want: "Hello how are you"},
		{name: "smoke 2", args: args{s: "What is the solution to this problem", k: 4}, want: "What is the solution"},
		{name: "smoke 3", args: args{s: "chopper is not a tanuki", k: 5}, want: "chopper is not a tanuki"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := truncateSentence(test.args.s, test.args.k); got != test.want {
				t.Errorf("truncateSentence() = %v, want = %v", got, test.want)
			}
		})
	}
}
