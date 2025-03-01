package golang

import (
	"reflect"
	"testing"
)

func TestApplyOperations(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 2, 2, 1, 1, 0}}, want: []int{1, 4, 2, 0, 0, 0}},
		{name: "smoke 2", args: args{nums: []int{0, 1}}, want: []int{1, 0}},
		{
			name: "test 8: wrong answer",
			args: args{nums: []int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272}},
			want: []int{1694, 399, 832, 1758, 412, 206, 272, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := applyOperations(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("applyOperations() = %v, want = %v", got, test.want)
			}
		})
	}
}
