package golang

import "testing"

func TestFindNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{12, 345, 2, 6, 7896}}, want: 2},
		{name: "smoke 2", args: args{nums: []int{555, 901, 482, 1771}}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findNumbers(test.args.nums); got != test.want {
				t.Errorf("findNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
