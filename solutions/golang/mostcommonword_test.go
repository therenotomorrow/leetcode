package golang

import "testing"

func TestMostCommonWord(t *testing.T) {
	t.Parallel()

	type args struct {
		paragraph string
		banned    []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smoke 1",
			args: args{paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.", banned: []string{"hit"}},
			want: "ball",
		},
		{
			name: "smoke 2",
			args: args{paragraph: "a.", banned: []string{}},
			want: "a",
		},
		{
			name: "test 36: wrong answer",
			args: args{paragraph: "a, a, a, a, b,b,b,c, c", banned: []string{"a"}}, //nolint:dupword
			want: "b",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := mostCommonWord(test.args.paragraph, test.args.banned); got != test.want {
				t.Errorf("mostCommonWord() = %v, want = %v", got, test.want)
			}
		})
	}
}
