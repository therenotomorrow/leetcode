package golang

import "testing"

func TestKthDistinct(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []string
		k   int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{arr: []string{"d", "b", "c", "b", "c", "a"}, k: 2}, want: "a"},
		{name: "smoke 2", args: args{arr: []string{"aaa", "aa", "a"}, k: 1}, want: "aaa"},
		{name: "smoke 3", args: args{arr: []string{"a", "b", "a"}, k: 3}, want: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := kthDistinct(test.args.arr, test.args.k); got != test.want {
				t.Errorf("kthDistinct() = %v, want = %v", got, test.want)
			}
		})
	}
}
