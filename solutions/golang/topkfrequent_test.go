package golang

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, want: []int{1, 2}},
		{name: "smoke 2", args: args{nums: []int{1}, k: 1}, want: []int{1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := topKFrequent(test.args.nums, test.args.k); !reflect.DeepEqual(got, test.want) {
				t.Errorf("topKFrequent() = %v, want = %v", got, test.want)
			}
		})
	}
}
