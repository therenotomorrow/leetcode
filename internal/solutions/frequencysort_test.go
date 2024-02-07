package solutions

import "testing"

func TestFrequencySort(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frequencySort(tt.args.s)

			for _, w := range tt.want {
				if got == w {
					return
				}
			}

			t.Errorf("frequencySort() = %v, want = %v", got, tt.want)
		})
	}
}
