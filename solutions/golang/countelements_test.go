package golang

import "testing"

func TestCountElements(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{arr: []int{1, 2, 3}}, want: 2},
		{name: "smoke 2", args: args{arr: []int{1, 1, 3, 3, 5, 5, 7, 7}}, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countElements(test.args.arr); got != test.want {
				t.Errorf("countElements() = %v, want = %v", got, test.want)
			}
		})
	}
}
