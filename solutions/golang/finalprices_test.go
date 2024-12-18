package golang

import (
	"reflect"
	"testing"
)

func TestFinalPrices(t *testing.T) {
	t.Parallel()

	type args struct {
		prices []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{prices: []int{8, 4, 6, 2, 3}}, want: []int{4, 2, 4, 2, 3}},
		{name: "smoke 2", args: args{prices: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
		{name: "smoke 3", args: args{prices: []int{10, 1, 1, 6}}, want: []int{9, 0, 1, 6}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := finalPrices(test.args.prices); !reflect.DeepEqual(got, test.want) {
				t.Errorf("finalPrices() = %v, want = %v", got, test.want)
			}
		})
	}
}
