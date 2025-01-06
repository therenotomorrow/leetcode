package golang

import (
	"reflect"
	"testing"
)

func TestMinOperations5(t *testing.T) {
	t.Parallel()

	type args struct {
		boxes string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{boxes: "110"}, want: []int{1, 1, 3}},
		{name: "smoke 2", args: args{boxes: "001011"}, want: []int{11, 8, 5, 4, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minOperations5(test.args.boxes); !reflect.DeepEqual(got, test.want) {
				t.Errorf("minOperations5() = %v, want = %v", got, test.want)
			}
		})
	}
}
