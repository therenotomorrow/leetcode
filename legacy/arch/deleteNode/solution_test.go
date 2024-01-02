package deleteNode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 7},
					},
				},
				key: 3,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 7},
					},
				},
				key: 0,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			args: args{
				root: nil,
				key:  0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				fmt.Println(tt.want)

				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNode2(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		list *ListNode
		want *ListNode
	}{
		{
			name: "1",
			args: args{node: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9},
				},
			}},
			list: &ListNode{Val: 4},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 9},
				},
			},
		},
		{
			name: "2",
			args: args{node: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 9},
			}},
			list: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5},
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: &ListNode{Val: 9},
				},
			},
		},
	}
	for _, tt := range tests {
		switch tt.name {
		case "1":
			tt.list.Next = tt.args.node

		case "2":
			tt.list.Next.Next = tt.args.node
		}

		t.Run(tt.name, func(t *testing.T) {
			deleteNode2(tt.args.node)
			if !reflect.DeepEqual(tt.list, tt.want) {
				t.Errorf("deleteNode2() = %v, want %v", tt.list, tt.want)
			}
		})
	}
}
