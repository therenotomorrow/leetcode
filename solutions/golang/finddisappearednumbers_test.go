package golang

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}}, want: []int{5, 6}},
		{name: "smoke 2", args: args{nums: []int{1, 1}}, want: []int{2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findDisappearedNumbers(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findDisappearedNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
