package golang

import (
	"reflect"
	"testing"
)

func TestGetMaximumXor(t *testing.T) {
	t.Parallel()

	type args struct {
		nums       []int
		maximumBit int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 1, 3}, maximumBit: 2}, want: []int{0, 3, 2, 3}},
		{name: "smoke 2", args: args{nums: []int{2, 3, 4, 7}, maximumBit: 3}, want: []int{5, 2, 6, 5}},
		{name: "smoke 3", args: args{nums: []int{0, 1, 2, 2, 5, 7}, maximumBit: 3}, want: []int{4, 3, 6, 4, 6, 7}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getMaximumXor(test.args.nums, test.args.maximumBit); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getMaximumXor() = %v, want = %v", got, test.want)
			}
		})
	}
}
