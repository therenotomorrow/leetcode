package golang

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	t.Parallel()

	type args struct {
		list1 []string
		list2 []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "smoke 1",
			args: args{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			want: []string{"Shogun"},
		},
		{
			name: "smoke 2",
			args: args{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"KFC", "Shogun", "Burger King"},
			},
			want: []string{"Shogun"},
		},
		{
			name: "smoke 3",
			args: args{
				list1: []string{"happy", "sad", "good"},
				list2: []string{"sad", "happy", "good"},
			},
			want: []string{"happy", "sad"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := findRestaurant(test.args.list1, test.args.list2)

			sort.Strings(got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("findRestaurant() = %v, want = %v", got, test.want)
			}
		})
	}
}
