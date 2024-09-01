package golang

import "testing"

func TestPathSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{113, 215, 221}}, want: 12},
		{name: "smoke 2", args: args{nums: []int{113, 221}}, want: 4},
		{name: "test 147: wrong answer", args: args{nums: []int{113, 229, 330, 466}}, want: 18},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := pathSum(test.args.nums); got != test.want {
				t.Errorf("pathSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
