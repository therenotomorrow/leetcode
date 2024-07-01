package golang

import "testing"

func TestMinCost1(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minCost1(test.args.colors, test.args.neededTime); got != test.want {
				t.Errorf("minCost1() = %v, want = %v", got, test.want)
			}
		})
	}
}
