package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)

			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) < len(got[j])
			})

			gotSlice := make([]string, 0)
			for _, group := range got {
				sort.Strings(group)
				gotSlice = append(gotSlice, group...)
			}

			sort.Slice(tt.want, func(i, j int) bool {
				return len(tt.want[i]) < len(tt.want[j])
			})

			wantSlice := make([]string, 0)
			for _, group := range tt.want {
				sort.Strings(group)

				wantSlice = append(wantSlice, group...)
			}

			if !reflect.DeepEqual(gotSlice, wantSlice) {
				t.Errorf("groupAnagrams() = %v, want = %v", got, tt.want)
			}
		})
	}
}
