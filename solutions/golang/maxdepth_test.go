package golang

import "testing"

func TestMaxDepth(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{s: "(1+(2*3)+((8)/4))+1"}, want: 3},
		{name: "smoke 2", args: args{s: "(1)+((2))+(((3)))"}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.s); got != tt.want {
				t.Errorf("maxDepth() = %v, want = %v", got, tt.want)
			}
		})
	}
}
