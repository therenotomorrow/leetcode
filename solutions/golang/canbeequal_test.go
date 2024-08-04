package golang

import "testing"

func TestCanBeEqual(t *testing.T) {
	t.Parallel()

	type args struct {
		target []int
		arr    []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{target: []int{1, 2, 3, 4}, arr: []int{2, 4, 1, 3}}, want: true},
		{name: "smoke 2", args: args{target: []int{7}, arr: []int{7}}, want: true},
		{name: "smoke 3", args: args{target: []int{3, 7, 9}, arr: []int{3, 7, 11}}, want: false},
		{name: "test 105: wrong answer", args: args{target: []int{1, 2, 3}, arr: []int{2, 2, 2}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := canBeEqual(test.args.target, test.args.arr); got != test.want {
				t.Errorf("canBeEqual() = %v, want = %v", got, test.want)
			}
		})
	}
}
