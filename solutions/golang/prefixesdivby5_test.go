package golang

import (
	"reflect"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	t.Parallel()

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
		{
			name: "test 2: wrong answer",
			args: args{nums: []int{0, 1, 1, 1, 1, 1}},
			want: []bool{true, false, false, false, true, false},
		},
		{
			name: "test 4: wrong answer",
			args: args{nums: []int{1, 1, 0, 0, 0, 1, 0, 0, 1}},
			want: []bool{false, false, false, false, false, false, false, false, false},
		},
		{
			name: "test 10: wrong answer",
			args: args{nums: []int{
				1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1,
				1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1,
				1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0,
			}},
			want: []bool{
				false, false, true, false, false, false, false, false, false, false,
				true, true, true, true, true, true, false, false, false, false, false,
				false, false, false, false, false, false, false, false, false, false,
				false, false, false, false, false, false, false, false, false, false,
				false, false, true, false, false, false, true, false, false, true,
				false, false, true, true, true, true, true, true, true, false, false,
				true, false, false, false, false, true, true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := prefixesDivBy5(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("prefixesDivBy5() = %v, want = %v", got, test.want)
			}
		})
	}
}
