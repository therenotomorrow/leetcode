package golang

import "testing"

func TestLeastInterval(t *testing.T) {
	t.Parallel()

	type args struct {
		tasks []byte
		n     int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2}, want: 8},
		{name: "smoke 2", args: args{tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'}, n: 1}, want: 6},
		{name: "smoke 3", args: args{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 3}, want: 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := leastInterval(test.args.tasks, test.args.n); got != test.want {
				t.Errorf("leastInterval() = %v, want = %v", got, test.want)
			}
		})
	}
}
