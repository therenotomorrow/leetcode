package golang

import "testing"

func TestGetLastMoment(t *testing.T) {
	type args struct {
		n     int
		left  []int
		right []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 4, left: []int{4, 3}, right: []int{0, 1}}, want: 4},
		{name: "smoke 2", args: args{n: 7, left: []int{}, right: []int{0, 1, 2, 3, 4, 5, 6, 7}}, want: 7},
		{name: "smoke 3", args: args{n: 7, left: []int{0, 1, 2, 3, 4, 5, 6, 7}, right: []int{}}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastMoment(tt.args.n, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("getLastMoment() = %v, want = %v", got, tt.want)
			}
		})
	}
}
