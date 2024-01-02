package solutions

import "testing"

func TestMakeEqual(t *testing.T) {
	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{words: []string{"abc", "aabc", "bc"}}, want: true},
		{name: "smoke 2", args: args{words: []string{"ab", "a"}}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeEqual(tt.args.words); got != tt.want {
				t.Errorf("makeEqual() = %v, want = %v", got, tt.want)
			}
		})
	}
}
