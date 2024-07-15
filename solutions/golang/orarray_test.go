package golang

import (
	"reflect"
	"testing"
)

func TestOrArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 3, 7, 15}}, want: []int{3, 7, 15}},
		{name: "smoke 2", args: args{nums: []int{8, 4, 2}}, want: []int{12, 6}},
		{name: "smoke 3", args: args{nums: []int{5, 4, 9, 11}}, want: []int{5, 13, 11}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := orArray(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("orArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
