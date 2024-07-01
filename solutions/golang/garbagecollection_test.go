package golang

import "testing"

func TestGarbageCollection(t *testing.T) {
	t.Parallel()

	type args struct {
		garbage []string
		travel  []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{garbage: []string{"G", "P", "GP", "GG"}, travel: []int{2, 4, 3}}, want: 21},
		{name: "smoke 2", args: args{garbage: []string{"MMM", "PGM", "GP"}, travel: []int{3, 10}}, want: 37},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := garbageCollection(test.args.garbage, test.args.travel); got != test.want {
				t.Errorf("garbageCollection() = %v, want = %v", got, test.want)
			}
		})
	}
}
