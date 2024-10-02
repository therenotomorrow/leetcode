package golang

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{arr: []int{40, 10, 20, 30}}, want: []int{4, 1, 2, 3}},
		{name: "smoke 2", args: args{arr: []int{100, 100, 100}}, want: []int{1, 1, 1}},
		{name: "smoke 3", args: args{arr: []int{37, 12, 28, 9, 100, 56, 80, 5, 12}}, want: []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := arrayRankTransform(test.args.arr); !reflect.DeepEqual(got, test.want) {
				t.Errorf("arrayRankTransform() = %v, want = %v", got, test.want)
			}
		})
	}
}
