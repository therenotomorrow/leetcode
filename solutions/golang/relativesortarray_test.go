package golang

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	t.Parallel()

	type args struct {
		arr1 []int
		arr2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, arr2: []int{2, 1, 4, 3, 9, 6}},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "smoke 2",
			args: args{arr1: []int{28, 6, 22, 8, 44, 17}, arr2: []int{22, 28, 8, 6}},
			want: []int{22, 28, 8, 6, 17, 44},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := relativeSortArray(test.args.arr1, test.args.arr2); !reflect.DeepEqual(got, test.want) {
				t.Errorf("relativeSortArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
