package golang

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "smoke 1", args: args{s: "ababcbacadefegdehijhklij"}, want: []int{9, 7, 8}},
		{name: "smoke 2", args: args{s: "eccbbbbdec"}, want: []int{10}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := partitionLabels(test.args.s); !reflect.DeepEqual(got, test.want) {
				t.Errorf("partitionLabels() = %v, want = %v", got, test.want)
			}
		})
	}
}
