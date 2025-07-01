package golang

import "testing"

func TestTriangleType(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{nums: []int{3, 3, 3}}, want: "equilateral"},
		{name: "smoke 2", args: args{nums: []int{3, 4, 5}}, want: "scalene"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := triangleType(test.args.nums); got != test.want {
				t.Errorf("triangleType() = %v, want = %v", got, test.want)
			}
		})
	}
}
