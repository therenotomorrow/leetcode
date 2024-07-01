package golang

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	t.Parallel()

	type args struct {
		rowIndex int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{rowIndex: 3}, want: []int{1, 3, 3, 1}},
		{name: "smoke 2", args: args{rowIndex: 0}, want: []int{1}},
		{name: "smoke 3", args: args{rowIndex: 1}, want: []int{1, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := getRow(test.args.rowIndex); !reflect.DeepEqual(got, test.want) {
				t.Errorf("getRow() = %v, want = %v", got, test.want)
			}
		})
	}
}
