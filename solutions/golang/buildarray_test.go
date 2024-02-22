package golang

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{target: []int{1, 3}, n: 3}, want: []string{"Push", "Push", "Pop", "Push"}},
		{name: "smoke 2", args: args{target: []int{1, 2, 3}, n: 3}, want: []string{"Push", "Push", "Push"}},
		{name: "smoke 3", args: args{target: []int{1, 2}, n: 4}, want: []string{"Push", "Push"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.target, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want = %v", got, tt.want)
			}
		})
	}
}
