package golang

import "testing"

func TestConnectSticks(t *testing.T) {
	type args struct {
		sticks []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{sticks: []int{2, 4, 3}}, want: 14},
		{name: "smoke 2", args: args{sticks: []int{1, 8, 3, 5}}, want: 30},
		{name: "smoke 3", args: args{sticks: []int{5}}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectSticks(tt.args.sticks); got != tt.want {
				t.Errorf("connectSticks() = %v, want = %v", got, tt.want)
			}
		})
	}
}
