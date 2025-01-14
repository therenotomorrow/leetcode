package golang

import (
	"reflect"
	"testing"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	t.Parallel()

	type args struct {
		A []int
		B []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}}, want: []int{0, 2, 3, 4}},
		{name: "smoke 2", args: args{A: []int{2, 3, 1}, B: []int{3, 1, 2}}, want: []int{0, 1, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findThePrefixCommonArray(test.args.A, test.args.B); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
