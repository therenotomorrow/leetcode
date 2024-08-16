package golang

import "testing"

func TestValidWordSquare(t *testing.T) {
	t.Parallel()

	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{words: []string{"abcd", "bnrt", "crmy", "dtye"}}, want: true},
		{name: "smoke 2", args: args{words: []string{"abcd", "bnrt", "crm", "dt"}}, want: true},
		{name: "smoke 3", args: args{words: []string{"ball", "area", "read", "lady"}}, want: false},
		{name: "test 4: wrong answer", args: args{words: []string{"ball", "asee", "let", "lep"}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := validWordSquare(test.args.words); got != test.want {
				t.Errorf("validWordSquare() = %v, want = %v", got, test.want)
			}
		})
	}
}
