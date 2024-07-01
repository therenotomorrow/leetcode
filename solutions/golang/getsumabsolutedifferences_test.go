package golang

import (
	"reflect"
	"testing"
)

func TestGetSumAbsoluteDifferences(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{2, 3, 5}}, want: []int{4, 3, 5}},
		{name: "smoke 2", args: args{nums: []int{1, 4, 6, 8, 10}}, want: []int{24, 15, 13, 15, 21}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getSumAbsoluteDifferences(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getSumAbsoluteDifferences() = %v, want = %v", got, test.want)
			}
		})
	}
}
