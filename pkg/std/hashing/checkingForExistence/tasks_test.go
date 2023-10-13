package checkingForExistence

import (
	"reflect"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{}}, want: []int{}},
		{args: args{nums: []int{1, 2, 3, 4, 5}}, want: []int{}},
		{args: args{nums: []int{1, 2, 0, 1, 5}}, want: []int{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
