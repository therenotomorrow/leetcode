package golang

import "testing"

func TestAverageWaitingTime(t *testing.T) {
	t.Parallel()

	type args struct {
		customers [][]int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "smoke 1", args: args{customers: [][]int{{1, 2}, {2, 5}, {4, 3}}}, want: 5},
		{name: "smoke 2", args: args{customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}}, want: 3.25},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := averageWaitingTime(test.args.customers); got != test.want {
				t.Errorf("averageWaitingTime() = %v, want = %v", got, test.want)
			}
		})
	}
}
