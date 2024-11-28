package golang

import (
	"reflect"
	"testing"
)

func TestMinimumAbsDifference(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "smoke 1", args: args{arr: []int{4, 2, 1, 3}}, want: [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{name: "smoke 2", args: args{arr: []int{1, 3, 6, 10, 15}}, want: [][]int{{1, 3}}},
		{
			name: "smoke 3",
			args: args{arr: []int{3, 8, -10, 23, 19, -4, -14, 27}},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
		{name: "test 9: wrong answer", args: args{arr: []int{40, 11, 26, 27, -20}}, want: [][]int{{26, 27}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := minimumAbsDifference(test.args.arr); !reflect.DeepEqual(got, test.want) {
				t.Errorf("minimumAbsDifference() = %v, want = %v", got, test.want)
			}
		})
	}
}
