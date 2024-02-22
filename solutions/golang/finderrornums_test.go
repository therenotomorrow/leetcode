package golang

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 2, 4}}, want: []int{2, 3}},
		{name: "smoke 2", args: args{nums: []int{1, 1}}, want: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want = %v", got, tt.want)
			}
		})
	}
}
