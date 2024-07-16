package MyLinkedList

import "testing"

func TestSolution(t *testing.T) {
	var obj MyLinkedList
	want := 0

	// ---- case 1
	obj = Constructor()

	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2) // linked list becomes 1->2->3

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj.Get(1); got != want {
			t.Errorf("obj.Get() = %v, want %v", got, want)
		}
	})

	obj.DeleteAtIndex(1) // now the linked list is 1->3

	want = 3
	t.Run("", func(t *testing.T) {
		if got := obj.Get(1); got != want {
			t.Errorf("obj.Get() = %v, want %v", got, want)
		}
	})

	// ---- case 2
	obj2 := Constructor()

	obj2.AddAtHead(2)
	obj2.DeleteAtIndex(1)
	obj2.AddAtHead(2)
	obj2.AddAtHead(7)
	obj2.AddAtHead(3)
	obj2.AddAtHead(2)
	obj2.AddAtHead(5)
	obj2.AddAtTail(5)

	want = 2
	t.Run("", func(t *testing.T) {
		if got := obj2.Get(5); got != want {
			t.Errorf("obj2.Get() = %v, want %v", got, want)
		}
	})

	obj2.DeleteAtIndex(6)
	obj2.DeleteAtIndex(4)

	// ---- case 3
	obj3 := Constructor()

	obj3.AddAtHead(4)

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj3.Get(1); got != want {
			t.Errorf("obj3.Get() = %v, want %v", got, want)
		}
	})

	obj3.AddAtHead(1)
	obj3.AddAtHead(5)
	obj3.DeleteAtIndex(3)
	obj3.AddAtHead(7)

	want = 4
	t.Run("", func(t *testing.T) {
		if got := obj3.Get(3); got != want {
			t.Errorf("obj3.Get() = %v, want %v", got, want)
		}
	})

	obj3.AddAtHead(1)
	obj3.DeleteAtIndex(4)
}
