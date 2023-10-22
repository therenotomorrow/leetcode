package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{numbers: []int{2, 7, 11, 15}, target: 9}, want: []int{1, 2}},
		{args: args{numbers: []int{2, 3, 4}, target: 6}, want: []int{1, 3}},
		{args: args{numbers: []int{-1, 0}, target: -1}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
