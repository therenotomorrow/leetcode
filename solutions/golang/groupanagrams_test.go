package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	t.Parallel()

	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "smoke 1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{name: "smoke 2", args: args{strs: []string{""}}, want: [][]string{{""}}},
		{name: "smoke 3", args: args{strs: []string{"a"}}, want: [][]string{{"a"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := groupAnagrams(test.args.strs)

			sort.SliceStable(got, func(i, j int) bool {
				return len(got[i]) < len(got[j])
			})

			gotSlice := make([]string, 0)

			for _, group := range got {
				sort.Strings(group)
				gotSlice = append(gotSlice, group...)
			}

			sort.SliceStable(test.want, func(i, j int) bool {
				return len(test.want[i]) < len(test.want[j])
			})

			wantSlice := make([]string, 0)

			for _, group := range test.want {
				sort.Strings(group)

				wantSlice = append(wantSlice, group...)
			}

			if !reflect.DeepEqual(gotSlice, wantSlice) {
				t.Errorf("groupAnagrams() = %v, want = %v", got, test.want)
			}
		})
	}
}
