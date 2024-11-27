package golang

import (
	"reflect"
	"testing"
)

func TestToArray(t *testing.T) {
	t.Parallel()

	type args struct {
		head *DoubleListNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "smoke 1",
			args: args{head: &DoubleListNode{
				Val: 1,
				Next: &DoubleListNode{
					Val: 2,
					Next: &DoubleListNode{
						Val: 3,
						Next: &DoubleListNode{
							Val: 4,
							Next: &DoubleListNode{
								Val: 3,
								Next: &DoubleListNode{
									Val:  2,
									Next: &DoubleListNode{Val: 1, Next: nil, Prev: nil},
									Prev: nil,
								},
								Prev: nil,
							},
							Prev: nil,
						},
						Prev: nil,
					},
					Prev: nil,
				},
				Prev: nil,
			}},
			want: []int{1, 2, 3, 4, 3, 2, 1},
		},
		{
			name: "smoke 2",
			args: args{head: &DoubleListNode{
				Val: 2,
				Next: &DoubleListNode{
					Val: 2,
					Next: &DoubleListNode{
						Val: 2,
						Next: &DoubleListNode{
							Val:  2,
							Next: &DoubleListNode{Val: 2, Next: nil, Prev: nil},
							Prev: nil,
						},
						Prev: nil,
					},
					Prev: nil,
				},
				Prev: nil,
			}},
			want: []int{2, 2, 2, 2, 2},
		},
		{
			name: "smoke 3",
			args: args{head: &DoubleListNode{
				Val: 3,
				Next: &DoubleListNode{
					Val: 2,
					Next: &DoubleListNode{
						Val: 3,
						Next: &DoubleListNode{
							Val: 2,
							Next: &DoubleListNode{
								Val:  3,
								Next: &DoubleListNode{Val: 2, Next: nil, Prev: nil},
								Prev: nil,
							},
							Prev: nil,
						},
						Prev: nil,
					},
					Prev: nil,
				},
				Prev: nil,
			}},
			want: []int{3, 2, 3, 2, 3, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := toArray(test.args.head); !reflect.DeepEqual(got, test.want) {
				t.Errorf("toArray() = %v, want = %v", got, test.want)
			}
		})
	}
}
