package golang

import "testing"

func TestCustomSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{order: "cba", s: "abcd"}, want: "cbad"},
		{name: "smoke 2", args: args{order: "cbafg", s: "abcd"}, want: "cbad"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("customSortString() = %v, want = %v", got, tt.want)
			}
		})
	}
}
