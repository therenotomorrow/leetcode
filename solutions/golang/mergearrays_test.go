package golang

import (
	"reflect"
	"testing"
)

func TestMergeArrays(t *testing.T) {
	t.Parallel()

	type args struct {
		nums1 [][]int
		nums2 [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "smoke 1",
			args: args{nums1: [][]int{{1, 2}, {2, 3}, {4, 5}}, nums2: [][]int{{1, 4}, {3, 2}, {4, 1}}},
			want: [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			name: "smoke 2",
			args: args{nums1: [][]int{{2, 4}, {3, 6}, {5, 5}}, nums2: [][]int{{1, 3}, {4, 3}}},
			want: [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := mergeArrays(test.args.nums1, test.args.nums2); !reflect.DeepEqual(got, test.want) {
				t.Errorf("mergeArrays() = %v, want = %v", got, test.want)
			}
		})
	}
}
