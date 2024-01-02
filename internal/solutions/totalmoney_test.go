package solutions

import "testing"

func TestTotalMoney(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalMoney(tt.args.n); got != tt.want {
				t.Errorf("totalMoney() = %v, want = %v", got, tt.want)
			}
		})
	}
}
