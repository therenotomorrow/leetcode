package golang

import "testing"

func TestMinStartValue(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{-3, 2, -3, 4, 2}}, want: 5},
		{name: "smoke 2", args: args{nums: []int{1, 2}}, want: 1},
		{name: "smoke 3", args: args{nums: []int{1, -2, -3}}, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minStartValue(test.args.nums); got != test.want {
				t.Errorf("minStartValue() = %v, want = %v", got, test.want)
			}
		})
	}
}
