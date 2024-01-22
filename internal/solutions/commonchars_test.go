package solutions

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want = %v", got, tt.want)
			}
		})
	}
}
