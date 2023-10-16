package solutions

import (
	"reflect"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{name: "smoke 1", args: args{nums: []int{0, 1, 1}}, want: []bool{true, false, false}},
		{name: "smoke 2", args: args{nums: []int{1, 1, 1}}, want: []bool{false, false, false}},
		{name: "test 2: wrong answer", args: args{nums: []int{0, 1, 1, 1, 1, 1}}, want: []bool{true, false, false, false, true, false}},
		{name: "test 4: wrong answer", args: args{nums: []int{1, 1, 0, 0, 0, 1, 0, 0, 1}}, want: []bool{false, false, false, false, false, false, false, false, false}},
		{name: "test 10: wrong answer", args: args{nums: []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}}, want: []bool{false, false, true, false, false, false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, true, false, false, false, false, true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want = %v", got, tt.want)
			}
		})
	}
}
