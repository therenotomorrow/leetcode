package golang

import "testing"

func TestDoesValidArrayExist(t *testing.T) {
	t.Parallel()

	type args struct {
		derived []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{derived: []int{1, 1, 0}}, want: true},
		{name: "smoke 2", args: args{derived: []int{1, 1}}, want: true},
		{name: "smoke 3", args: args{derived: []int{1, 0}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := doesValidArrayExist(test.args.derived); got != test.want {
				t.Errorf("doesValidArrayExist() = %v, want = %v", got, test.want)
			}
		})
	}
}
