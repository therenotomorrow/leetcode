package solutions

import "testing"

func TestTitleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{columnTitle: "A"}, want: 1},
		{name: "smoke 2", args: args{columnTitle: "AB"}, want: 28},
		{name: "smoke 3", args: args{columnTitle: "ZY"}, want: 701},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want = %v", got, tt.want)
			}
		})
	}
}
