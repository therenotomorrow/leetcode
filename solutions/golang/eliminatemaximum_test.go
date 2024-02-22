package golang

import "testing"

func TestEliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{dist: []int{1, 3, 4}, speed: []int{1, 1, 1}}, want: 3},
		{name: "smoke 2", args: args{dist: []int{1, 1, 2, 3}, speed: []int{1, 1, 1, 1}}, want: 1},
		{name: "smoke 3", args: args{dist: []int{3, 2, 4}, speed: []int{5, 3, 2}}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMaximum(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want = %v", got, tt.want)
			}
		})
	}
}
