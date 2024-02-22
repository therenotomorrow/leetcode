package golang

import "testing"

func TestMinCost1(t *testing.T) {
	type args struct {
		colors     string
		neededTime []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{colors: "abaac", neededTime: []int{1, 2, 3, 4, 5}}, want: 3},
		{name: "smoke 2", args: args{colors: "abc", neededTime: []int{1, 2, 3}}, want: 0},
		{name: "smoke 3", args: args{colors: "aabaa", neededTime: []int{1, 2, 3, 4, 1}}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost1(tt.args.colors, tt.args.neededTime); got != tt.want {
				t.Errorf("minCost1() = %v, want = %v", got, tt.want)
			}
		})
	}
}
