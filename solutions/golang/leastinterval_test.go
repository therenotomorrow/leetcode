package golang

import "testing"

func TestLeastInterval(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want = %v", got, tt.want)
			}
		})
	}
}
