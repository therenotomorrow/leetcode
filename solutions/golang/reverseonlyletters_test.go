package golang

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want = %v", got, tt.want)
			}
		})
	}
}
