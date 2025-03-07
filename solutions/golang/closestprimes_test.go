package golang

import (
	"reflect"
	"testing"
)

func TestClosestPrimes(t *testing.T) {
	t.Parallel()

	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{left: 10, right: 19}, want: []int{11, 13}},
		{name: "smoke 2", args: args{left: 4, right: 6}, want: []int{-1, -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := closestPrimes(test.args.left, test.args.right); !reflect.DeepEqual(got, test.want) {
				t.Errorf("closestPrimes() = %v, want = %v", got, test.want)
			}
		})
	}
}
