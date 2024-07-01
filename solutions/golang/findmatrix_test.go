package golang

import (
	"reflect"
	"testing"
)

func TestFindMatrix(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 4, 1, 2, 3, 1}}, want: [][]int{{1, 2, 3, 4}, {1, 3}, {1}}},
		{name: "smoke 2", args: args{nums: []int{1, 2, 3, 4}}, want: [][]int{{1, 2, 3, 4}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findMatrix(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findMatrix() = %v, want = %v", got, test.want)
			}
		})
	}
}
