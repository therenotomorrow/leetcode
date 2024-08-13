package golang

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	t.Parallel()

	type args struct {
		candidates []int
		target     int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
			want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name: "smoke 2",
			args: args{candidates: []int{2, 5, 2, 1, 2}, target: 5},
			want: [][]int{{1, 2, 2}, {5}},
		},
		{
			name: "test 121: wrong answer",
			args: args{candidates: []int{3, 1, 3, 5, 1, 1}, target: 8},
			want: [][]int{{1, 1, 1, 5}, {1, 1, 3, 3}, {3, 5}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := combinationSum2(test.args.candidates, test.args.target); !reflect.DeepEqual(got, test.want) {
				t.Errorf("combinationSum2() = %v, want = %v", got, test.want)
			}
		})
	}
}
