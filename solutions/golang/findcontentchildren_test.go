package golang

import "testing"

func TestFindContentChildren(t *testing.T) {
	t.Parallel()

	type args struct {
		g []int
		s []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{g: []int{1, 2, 3}, s: []int{1, 1}}, want: 1},
		{name: "smoke 2", args: args{g: []int{1, 2}, s: []int{1, 2, 3}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findContentChildren(test.args.g, test.args.s); got != test.want {
				t.Errorf("findContentChildren() = %v, want = %v", got, test.want)
			}
		})
	}
}
