package golang

import "testing"

func TestFindKthNumber(t *testing.T) {
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
		{name: "smoke 1", args: args{n: 13, k: 2}, want: 10},
		{name: "smoke 2", args: args{n: 1, k: 1}, want: 1},
		{name: "test 39: runtime", args: args{n: 804289384, k: 42641503}, want: 138377349},
		{name: "test 42: runtime", args: args{n: 719885387, k: 209989719}, want: 288990744},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findKthNumber(test.args.n, test.args.k); got != test.want {
				t.Errorf("findKthNumber() = %v, want = %v", got, test.want)
			}
		})
	}
}
