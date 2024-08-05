package golang

import "testing"

func TestCountSeniors(t *testing.T) {
	t.Parallel()

	type args struct {
		details []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}},
			want: 2,
		},
		{
			name: "smoke 2",
			args: args{details: []string{"1313579440F2036", "2921522980M5644"}},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countSeniors(test.args.details); got != test.want {
				t.Errorf("countSeniors() = %v, want = %v", got, test.want)
			}
		})
	}
}
