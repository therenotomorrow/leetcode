package golang

import "testing"

func TestCheckInclusion(t *testing.T) {
	t.Parallel()

	type args struct {
		s1 string
		s2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{s1: "ab", s2: "eidbaooo"}, want: true},
		{name: "smoke 2", args: args{s1: "ab", s2: "eidboaoo"}, want: false},
		{name: "test 81: wrong answer", args: args{s1: "adc", s2: "dcda"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := checkInclusion(test.args.s1, test.args.s2); got != test.want {
				t.Errorf("checkInclusion() = %v, want = %v", got, test.want)
			}
		})
	}
}
