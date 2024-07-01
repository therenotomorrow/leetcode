package golang

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{s: "ab-cd"}, want: "dc-ba"},
		{name: "smoke 2", args: args{s: "a-bC-dEf-ghIj"}, want: "j-Ih-gfE-dCba"},
		{name: "smoke 3", args: args{s: "Test1ng-Leet=code-Q!"}, want: "Qedo1ct-eeLg=ntse-T!"},
		{name: "test 4: runtime", args: args{s: "-"}, want: "-"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := reverseOnlyLetters(test.args.s); got != test.want {
				t.Errorf("reverseOnlyLetters() = %v, want = %v", got, test.want)
			}
		})
	}
}
