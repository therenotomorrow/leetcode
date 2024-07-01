package golang

import "testing"

func TestFindSpecialInteger(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{1, 2, 2, 6, 6, 6, 6, 7, 10}}, want: 6},
		{name: "smoke 2", args: args{arr: []int{1, 1}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findSpecialInteger(test.args.arr); got != test.want {
				t.Errorf("findSpecialInteger() = %v, want = %v", got, test.want)
			}
		})
	}
}
