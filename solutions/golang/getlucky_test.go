package golang

import "testing"

func TestGetLucky(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "iiii", k: 1}, want: 36},
		{name: "smoke 2", args: args{s: "leetcode", k: 2}, want: 6},
		{name: "smoke 3", args: args{s: "zbax", k: 2}, want: 8},
		{name: "test 192: wrong answer", args: args{s: "iaozzbyqzwbpurzze", k: 2}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getLucky(test.args.s, test.args.k); got != test.want {
				t.Errorf("getLucky() = %v, want = %v", got, test.want)
			}
		})
	}
}
