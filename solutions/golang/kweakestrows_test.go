package golang

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{mat: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			}, k: 3},
			want: []int{2, 0, 3},
		},
		{
			name: "smoke 2",
			args: args{mat: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			}, k: 2},
			want: []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want = %v", got, tt.want)
			}
		})
	}
}
