package golang

import "testing"

func TestCountMatches(t *testing.T) {
	t.Parallel()

	type args struct {
		items     [][]string
		ruleKey   string
		ruleValue string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smoke 1",
			args: args{
				items: [][]string{
					{"phone", "blue", "pixel"},
					{"computer", "silver", "lenovo"},
					{"phone", "gold", "iphone"},
				},
				ruleKey:   "color",
				ruleValue: "silver",
			},
			want: 1,
		},
		{
			name: "smoke 2",
			args: args{
				items: [][]string{
					{"phone", "blue", "pixel"},
					{"computer", "silver", "phone"},
					{"phone", "gold", "iphone"},
				},
				ruleKey:   "type",
				ruleValue: "phone",
			},
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := countMatches(test.args.items, test.args.ruleKey, test.args.ruleValue); got != test.want {
				t.Errorf("countMatches() = %v, want = %v", got, test.want)
			}
		})
	}
}
