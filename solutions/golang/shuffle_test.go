package golang

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		n    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 5, 1, 3, 4, 7}, n: 3}, want: []int{2, 3, 5, 4, 1, 7}},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4, 4, 3, 2, 1}, n: 4}, want: []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{name: "smoke 3", args: args{nums: []int{1, 1, 2, 2}, n: 2}, want: []int{1, 2, 1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := shuffle(test.args.nums, test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("shuffle() = %v, want = %v", got, test.want)
			}
		})
	}
}
