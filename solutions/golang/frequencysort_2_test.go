package golang

import (
	"reflect"
	"testing"
)

func TestFrequencySort2(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{nums: []int{1, 1, 2, 2, 2, 3}},
			want: []int{3, 1, 1, 2, 2, 2},
		},
		{
			name: "smoke 2",
			args: args{nums: []int{2, 3, 1, 3, 2}},
			want: []int{1, 3, 3, 2, 2},
		},
		{
			name: "smoke 3",
			args: args{nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}},
			want: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := frequencySort2(test.args.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("frequencySort2() = %v, want = %v", got, test.want)
			}
		})
	}
}
