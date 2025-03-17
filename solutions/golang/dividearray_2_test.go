package golang

import "testing"

func TestDivideArray2(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{nums: []int{3, 2, 3, 2, 2, 2}}, want: true},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := divideArray2(test.args.nums); got != test.want {
				t.Errorf("divideArray2() = %v, want = %v", got, test.want)
			}
		})
	}
}
