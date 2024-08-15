package golang

import "testing"

func TestLemonadeChange(t *testing.T) {
	t.Parallel()

	type args struct {
		bills []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "smoke 1", args: args{bills: []int{5, 5, 5, 10, 20}}, want: true},
		{name: "smoke 2", args: args{bills: []int{5, 5, 10, 10, 20}}, want: false},
		{name: "test 24: wrong answer", args: args{bills: []int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20}}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := lemonadeChange(test.args.bills); got != test.want {
				t.Errorf("lemonadeChange() = %v, want = %v", got, test.want)
			}
		})
	}
}
