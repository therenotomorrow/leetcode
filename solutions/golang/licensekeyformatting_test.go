package golang

import "testing"

func TestLicenseKeyFormatting(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "5F3Z-2e-9-w", k: 4}, want: "5F3Z-2E9W"},
		{name: "smoke 2", args: args{s: "2-5g-3-J", k: 2}, want: "2-5G-3J"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := licenseKeyFormatting(test.args.s, test.args.k); got != test.want {
				t.Errorf("licenseKeyFormatting() = %v, want = %v", got, test.want)
			}
		})
	}
}
