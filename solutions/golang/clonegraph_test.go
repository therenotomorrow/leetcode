package golang

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	t.Parallel()

	type args struct {
		node *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: Smoke1, args: args{node: &Node{Val: 1, Children: nil}},
			want: &Node{Val: 1, Children: nil},
		},
		{
			name: Smoke2, args: args{node: &Node{Val: 1, Children: nil}},
			want: &Node{Val: 1, Children: nil},
		},
		{
			name: Smoke3, args: args{node: nil},
			want: nil,
		},
	}

	for _, test := range tests {
		switch test.name {
		case Smoke1:
			node1 := &Node{Val: 1, Children: nil}
			node2 := &Node{Val: 2, Children: nil}
			node3 := &Node{Val: 3, Children: nil}
			node4 := &Node{Val: 4, Children: nil}

			node1.Children = []*Node{node2, node4}
			node2.Children = []*Node{node1, node3}
			node3.Children = []*Node{node2, node4}
			node4.Children = []*Node{node1, node3}

			test.args.node = node1
			test.want = node1
		case Smoke2:
			node1 := &Node{Val: 1, Children: nil}

			test.args.node = node1
			test.want = node1
		case Smoke3:
		}

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := cloneGraph(test.args.node); !reflect.DeepEqual(got, test.want) {
				t.Errorf("cloneGraph() = %v, want = %v", got, test.want)
			}
		})
	}
}
