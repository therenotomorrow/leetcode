package areOccurrencesEqual

import "testing"

func Test_areOccurrencesEqual(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "abacbc"}, want: true},
		{args: args{s: "aaabb"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areOccurrencesEqual(tt.args.s); got != tt.want {
				t.Errorf("areOccurrencesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
