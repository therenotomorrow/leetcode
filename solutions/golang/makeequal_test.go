package golang

import "testing"

func TestMakeEqual(t *testing.T) {
	t.Parallel()

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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := makeEqual(test.args.words); got != test.want {
				t.Errorf("makeEqual() = %v, want = %v", got, test.want)
			}
		})
	}
}
