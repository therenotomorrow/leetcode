package golang

import (
	"reflect"
	"testing"
)

func TestLuckyNumbers(t *testing.T) {
	t.Parallel()

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{matrix: [][]int{
				{3, 7, 8},
				{9, 11, 13},
				{15, 16, 17},
			}},
			want: []int{15},
		},
		{
			name: "smoke 2",
			args: args{matrix: [][]int{
				{1, 10, 4, 2},
				{9, 3, 8, 7},
				{15, 16, 17, 12},
			}},
			want: []int{12},
		},
		{
			name: "smoke 3",
			args: args{matrix: [][]int{
				{7, 8},
				{1, 2},
			}},
			want: []int{7},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := luckyNumbers(test.args.matrix); !reflect.DeepEqual(got, test.want) {
				t.Errorf("luckyNumbers() = %v, want = %v", got, test.want)
			}
		})
	}
}
