package golang

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{nums: []string{"01", "10"}}, want: []string{"00", "11"}},
		{name: "smoke 2", args: args{nums: []string{"00", "01"}}, want: []string{"11", "10"}},
		{name: "smoke 3", args: args{nums: []string{"111", "011", "001"}}, want: []string{"101", "000", "010", "100", "110"}},
		{name: "test 4: runtime", args: args{nums: []string{"1"}}, want: []string{"0"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := findDifferentBinaryString(test.args.nums)

			// check any of possible solutions
			for _, w := range test.want {
				if got == w {
					return
				}
			}

			t.Errorf("findDifferentBinaryString() = %v, want = %v", got, test.want)
		})
	}
}
