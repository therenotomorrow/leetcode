package golang

import (
	"reflect"
	"testing"
)

func TestTransformArray(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{arr: []int{6, 2, 3, 4}}, want: []int{6, 3, 3, 4}},
		{name: "smoke 2", args: args{arr: []int{1, 6, 3, 4, 3, 5}}, want: []int{1, 4, 4, 4, 4, 5}},
		{
			name: "test 12: wrong answer",
			args: args{arr: []int{6, 5, 8, 6, 7, 7, 3, 9, 8, 8, 3, 1, 2, 9, 8, 3}},
			want: []int{6, 6, 7, 7, 7, 7, 7, 8, 8, 8, 3, 2, 2, 8, 8, 3},
		},
		{
			name: "test 16: wrong answer",
			args: args{arr: []int{2, 1, 2, 1, 1, 2, 2, 1}},
			want: []int{2, 2, 1, 1, 1, 2, 2, 1},
		},
		{
			name: "test 19: wrong answer",
			args: args{arr: []int{1, 5, 4, 7, 9, 2, 5, 1, 2, 5, 8, 8, 3, 8, 4, 4, 6, 3}},
			want: []int{1, 4, 5, 7, 7, 4, 3, 2, 2, 5, 8, 8, 6, 5, 4, 4, 4, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := transformArray(test.args.arr); !reflect.DeepEqual(got, test.want) {
				t.Errorf("transformArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
