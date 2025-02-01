package golang

import "testing"

func TestIsArraySpecial(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{1}}, want: true},
		{name: "smoke 2", args: args{nums: []int{2, 1, 4}}, want: true},
		{name: "smoke 3", args: args{nums: []int{4, 3, 1, 6}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := isArraySpecial(test.args.nums); got != test.want {
				t.Errorf("isArraySpecial() = %v, want = %v", got, test.want)
			}
		})
	}
}
