package golang

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 1}}, want: []int{1, 2, 1, 1, 2, 1}},
		{name: "smoke 2", args: args{nums: []int{1, 3, 2, 1}}, want: []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getConcatenation(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getConcatenation() = %v, want = %v", got, test.want)
			}
		})
	}
}
