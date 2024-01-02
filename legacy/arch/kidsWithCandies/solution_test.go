package kidsWithCandies

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{args: args{candies: []int{2, 3, 5, 1, 3}, extraCandies: 3}, want: []bool{true, true, true, false, true}},
		{args: args{candies: []int{4, 2, 1, 1, 2}, extraCandies: 1}, want: []bool{true, false, false, false, false}},
		{args: args{candies: []int{12, 1, 12}, extraCandies: 10}, want: []bool{true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
