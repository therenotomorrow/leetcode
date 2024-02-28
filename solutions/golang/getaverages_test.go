package golang

import (
	"reflect"
	"testing"
)

func TestGetAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{nums: []int{7, 4, 3, 9, 1, 8, 5, 2, 6}, k: 3},
			want: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{name: "smoke 2", args: args{nums: []int{100000}, k: 0}, want: []int{100000}},
		{name: "smoke 3", args: args{nums: []int{8}, k: 100000}, want: []int{-1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAverages(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAverages() = %v, want = %v", got, tt.want)
			}
		})
	}
}
