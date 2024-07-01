package golang

import (
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := findRelativeRanks(test.args.score); !reflect.DeepEqual(got, test.want) {
				t.Errorf("findRelativeRanks() = %v, want = %v", got, test.want)
			}
		})
	}
}
