package golang

import (
	"reflect"
	"testing"
)

func TestCheckArithmeticSubarrays(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		l    []int
		r    []int
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "smoke 1",
			args: args{
				nums: []int{4, 6, 5, 9, 3, 7},
				l:    []int{0, 0, 2},
				r:    []int{2, 3, 5},
			},
			want: []bool{true, false, true},
		},
		{
			name: "smoke 2",
			args: args{
				nums: []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10},
				l:    []int{0, 1, 6, 4, 8, 7},
				r:    []int{4, 4, 9, 7, 9, 10},
			},
			want: []bool{false, true, false, false, true, true},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := checkArithmeticSubarrays(test.args.nums, test.args.l, test.args.r); !reflect.DeepEqual(got, test.want) {
				t.Errorf("checkArithmeticSubarrays() = %v, want = %v", got, test.want)
			}
		})
	}
}
