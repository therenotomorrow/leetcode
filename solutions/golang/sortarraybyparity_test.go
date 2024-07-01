package golang

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{3, 1, 2, 4}}, want: []int{2, 4, 3, 1}},
		{name: "smoke 2", args: args{nums: []int{0}}, want: []int{0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := sortArrayByParity(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("sortArrayByParity() = %v, want = %v", got, test.want)
			}
		})
	}
}
