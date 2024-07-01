package golang

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 3, 4}}, want: []int{1, 3, 6, 10}},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1, 1, 1}}, want: []int{1, 2, 3, 4, 5}},
		{name: "smoke 3", args: args{nums: []int{3, 1, 2, 10, 1}}, want: []int{3, 4, 6, 16, 17}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := runningSum(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("runningSum() = %v, want = %v", got, test.want)
			}
		})
	}
}
