package golang

import "testing"

func TestValidSquare(t *testing.T) {
	t.Parallel()

	type args struct {
		p1 []int
		p2 []int
		p3 []int
		p4 []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{p1: []int{0, 0}, p2: []int{1, 1}, p3: []int{1, 0}, p4: []int{0, 1}}, want: true},
		{name: "smoke 2", args: args{p1: []int{0, 0}, p2: []int{1, 1}, p3: []int{1, 0}, p4: []int{0, 12}}, want: false},
		{name: "smoke 3", args: args{p1: []int{1, 0}, p2: []int{-1, 0}, p3: []int{0, 1}, p4: []int{0, -1}}, want: true},
		{
			name: "test 255: wrong answer",
			args: args{p1: []int{0, 0}, p2: []int{0, 0}, p3: []int{0, 0}, p4: []int{0, 0}},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := validSquare(test.args.p1, test.args.p2, test.args.p3, test.args.p4); got != test.want {
				t.Errorf("validSquare() = %v, want = %v", got, test.want)
			}
		})
	}
}
