package golang

import "testing"

func TestFirstCompleteIndex(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
		mat [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{arr: []int{1, 3, 4, 2}, mat: [][]int{{1, 4}, {2, 3}}},
			want: 2,
		},
		{
			name: "smoke 2",
			args: args{arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9}, mat: [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := firstCompleteIndex(test.args.arr, test.args.mat); got != test.want {
				t.Errorf("firstCompleteIndex() = %v, want = %v", got, test.want)
			}
		})
	}
}
