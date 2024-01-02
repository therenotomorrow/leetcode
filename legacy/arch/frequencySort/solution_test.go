package frequencySort

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{args: args{s: "tree"}, want: []string{"eert", "eetr"}},
		{args: args{s: "cccaaa"}, want: []string{"aaaccc", "cccaaa"}},
		{args: args{s: "Aabb"}, want: []string{"bbAa", "bbaA"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frequencySort(tt.args.s)

			for _, w := range tt.want {
				if got == w {
					return
				}
			}

			t.Errorf("frequencySort() = %v, want %v", got, tt.want)
		})
	}
}
