package golang

import "testing"

func TestNumberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{bank: []string{"011001", "000000", "010100", "001000"}}, want: 8},
		{name: "smoke 2", args: args{bank: []string{"000", "111", "000"}}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeams(tt.args.bank); got != tt.want {
				t.Errorf("numberOfBeams() = %v, want = %v", got, tt.want)
			}
		})
	}
}
