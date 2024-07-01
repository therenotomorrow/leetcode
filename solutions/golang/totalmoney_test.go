package golang

import "testing"

func TestTotalMoney(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 4}, want: 10},
		{name: "smoke 2", args: args{n: 10}, want: 37},
		{name: "smoke 3", args: args{n: 20}, want: 96},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := totalMoney(test.args.n); got != test.want {
				t.Errorf("totalMoney() = %v, want = %v", got, test.want)
			}
		})
	}
}
