package solutions

import "testing"

func TestSortVowels(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "lEetcOde"}, want: "lEOtcede"},
		{name: "smoke 2", args: args{s: "lYmpH"}, want: "lYmpH"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortVowels(tt.args.s); got != tt.want {
				t.Errorf("sortVowels() = %v, want = %v", got, tt.want)
			}
		})
	}
}
