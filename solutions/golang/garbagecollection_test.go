package golang

import "testing"

func TestGarbageCollection(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := garbageCollection(tt.args.garbage, tt.args.travel); got != tt.want {
				t.Errorf("garbageCollection() = %v, want = %v", got, tt.want)
			}
		})
	}
}
