package golang

import (
	"reflect"
	"testing"
)

func TestBuildArray1(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := buildArray1(test.args.target, test.args.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("buildArray1() = %v, want = %v", got, test.want)
			}
		})
	}
}
