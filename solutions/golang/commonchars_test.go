package golang

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{words: []string{"bella", "label", "roller"}}, want: []string{"e", "l", "l"}},
		{name: "smoke 2", args: args{words: []string{"cool", "lock", "cook"}}, want: []string{"c", "o"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := commonChars(test.args.words); !reflect.DeepEqual(got, test.want) {
				t.Errorf("commonChars() = %v, want = %v", got, test.want)
			}
		})
	}
}
