package prefixSum

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{args: args{nums: []int{1, 6, 3, 2, 7, 2}, queries: [][]int{{0, 3}, {2, 5}, {2, 4}}, limit: 13}, want: []bool{true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerQueries(tt.args.nums, tt.args.queries, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
