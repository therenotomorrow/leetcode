package golang

import (
	"reflect"
	"testing"
)

func TestEvenOddBit(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{n: 17}, want: []int{2, 0}},
		{name: "smoke 2", args: args{n: 2}, want: []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := evenOddBit(test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("evenOddBit() = %v, want = %v", got, test.want)
			}
		})
	}
}
