package golang

import "testing"

func TestMajorityElement(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{nums: []int{3, 2, 3}}, want: 3},
		{name: "smoke 2", args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}}, want: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := majorityElement(test.args.nums); got != test.want {
				t.Errorf("majorityElement() = %v, want = %v", got, test.want)
			}
		})
	}
}
