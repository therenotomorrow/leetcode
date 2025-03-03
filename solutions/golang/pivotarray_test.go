package golang

import (
	"reflect"
	"testing"
)

func TestPivotArray(t *testing.T) {
	t.Parallel()

	type args struct {
		nums  []int
		pivot int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{nums: []int{9, 12, 5, 10, 14, 3, 10}, pivot: 10},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name: "smoke 2",
			args: args{nums: []int{-3, 4, 3, 2}, pivot: 2},
			want: []int{-3, 2, 4, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := pivotArray(test.args.nums, test.args.pivot); !reflect.DeepEqual(got, test.want) {
				t.Errorf("pivotArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
