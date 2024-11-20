package golang

import (
	"reflect"
	"testing"
)

func TestNumberOfPairs(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 2, 1, 3, 2, 2}}, want: []int{3, 1}},
		{name: "smoke 2", args: args{nums: []int{1, 1}}, want: []int{1, 0}},
		{name: "smoke 3", args: args{nums: []int{0}}, want: []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := numberOfPairs(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("numberOfPairs() = %v, want = %v", got, test.want)
			}
		})
	}
}
