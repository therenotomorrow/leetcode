package golang

import "testing"

func TestPassThePillow(t *testing.T) {
	t.Parallel()

	type args struct {
		n    int
		time int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 4, time: 5}, want: 2},
		{name: "smoke 2", args: args{n: 3, time: 2}, want: 3},
		{name: "test 47: wrong answer", args: args{n: 18, time: 38}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := passThePillow(test.args.n, test.args.time); got != test.want {
				t.Errorf("passThePillow() = %v, want = %v", got, test.want)
			}
		})
	}
}
