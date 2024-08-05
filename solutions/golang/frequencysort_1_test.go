package golang

import "testing"

func TestFrequencySort1(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "smoke 1", args: args{s: "tree"}, want: []string{"eert", "eetr"}},
		{name: "smoke 2", args: args{s: "cccaaa"}, want: []string{"aaaccc", "cccaaa"}},
		{name: "smoke 3", args: args{s: "Aabb"}, want: []string{"bbAa", "bbaA"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := frequencySort1(test.args.s)

			for _, w := range test.want {
				if got == w {
					return
				}
			}

			t.Errorf("frequencySort1() = %v, want = %v", got, test.want)
		})
	}
}
