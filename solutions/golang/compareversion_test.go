package golang

import "testing"

func TestCompareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{version1: "1.2", version2: "1.10"}, want: -1},
		{name: "smoke 2", args: args{version1: "1.01", version2: "1.001"}, want: 0},
		{name: "smoke 3", args: args{version1: "1.0", version2: "1.0.0.0"}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want = %v", got, tt.want)
			}
		})
	}
}
