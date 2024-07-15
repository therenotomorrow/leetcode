package golang

import (
	"reflect"
	"testing"
)

func TestBuildArray2(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{0, 2, 1, 5, 3, 4}}, want: []int{0, 1, 2, 4, 5, 3}},
		{name: "smoke 2", args: args{nums: []int{5, 0, 1, 2, 3, 4}}, want: []int{4, 5, 0, 1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := buildArray2(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("buildArray2() = %v, want = %v", got, test.want)
			}
		})
	}
}
