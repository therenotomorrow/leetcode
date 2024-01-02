package solutions

import "testing"

func TestConvertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{columnNumber: 1}, want: "A"},
		{name: "smoke 2", args: args{columnNumber: 28}, want: "AB"},
		{name: "smoke 3", args: args{columnNumber: 701}, want: "ZY"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want = %v", got, tt.want)
			}
		})
	}
}
