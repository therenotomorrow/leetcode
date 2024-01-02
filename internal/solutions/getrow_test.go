package solutions

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want = %v", got, tt.want)
			}
		})
	}
}
