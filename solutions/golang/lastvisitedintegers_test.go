package golang

import (
	"reflect"
	"testing"
)

func TestLastVisitedIntegers(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, -1, -1, -1}}, want: []int{2, 1, -1}},
		{name: "smoke 2", args: args{nums: []int{1, -1, 2, -1, -1}}, want: []int{1, 2, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lastVisitedIntegers(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("lastVisitedIntegers() = %v, want = %v", got, test.want)
			}
		})
	}
}
