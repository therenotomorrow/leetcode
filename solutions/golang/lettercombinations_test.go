package golang

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	t.Parallel()

	type args struct {
		digits string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{digits: "23"}, want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "smoke 2", args: args{digits: ""}, want: []string{}},
		{name: "smoke 3", args: args{digits: "2"}, want: []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := letterCombinations(test.args.digits); !reflect.DeepEqual(got, test.want) {
				t.Errorf("letterCombinations() = %v, want = %v", got, test.want)
			}
		})
	}
}
