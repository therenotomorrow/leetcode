package golang

import (
	"testing"
)

func TestFindElementsSmoke1(t *testing.T) {
	t.Parallel()

	obj := FindElementsConstructor(&TreeNode{
		Val:   -1,
		Left:  nil,
		Right: &TreeNode{Val: -1, Left: nil, Right: nil},
	})

	if got := obj.Find(1); got != false {
		t.Errorf("obj.Find() = %v, want = %v", got, false)
	}

	if got := obj.Find(2); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}
}

func TestFindElementsSmoke2(t *testing.T) {
	t.Parallel()

	obj := FindElementsConstructor(&TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -1,
			Left:  &TreeNode{Val: -1, Left: nil, Right: nil},
			Right: &TreeNode{Val: -1, Left: nil, Right: nil},
		},
		Right: &TreeNode{Val: -1, Left: nil, Right: nil},
	})

	if got := obj.Find(1); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}

	if got := obj.Find(3); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}

	if got := obj.Find(5); got != false {
		t.Errorf("obj.Find() = %v, want = %v", got, false)
	}
}

func TestFindElementsSmoke3(t *testing.T) {
	t.Parallel()

	obj := FindElementsConstructor(&TreeNode{
		Val:  -1,
		Left: nil,
		Right: &TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val:   -1,
				Left:  &TreeNode{Val: -1, Left: nil, Right: nil},
				Right: nil,
			},
			Right: nil,
		},
	})

	if got := obj.Find(2); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}

	if got := obj.Find(3); got != false {
		t.Errorf("obj.Find() = %v, want = %v", got, false)
	}

	if got := obj.Find(4); got != false {
		t.Errorf("obj.Find() = %v, want = %v", got, false)
	}

	if got := obj.Find(5); got != true {
		t.Errorf("obj.Find() = %v, want = %v", got, true)
	}
}
