package golang

import (
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{score: []int{5, 4, 3, 2, 1}},
			want: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			name: "smoke 2",
			args: args{score: []int{10, 3, 8, 9, 4}},
			want: []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanks(tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want = %v", got, tt.want)
			}
		})
	}
}
