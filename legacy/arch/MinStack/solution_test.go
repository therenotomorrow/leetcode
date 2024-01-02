package MinStack

import "testing"

func TestSolution(t *testing.T) {
	var want int

	// ---- case 1
	obj := Constructor()

	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	want = -3
	t.Run("", func(t *testing.T) {
		if got := obj.GetMin(); got != want {
			t.Errorf(" obj.GetMin() = %v, want %v", got, want)
		}
	})

	obj.Pop()

	want = 0
	t.Run("", func(t *testing.T) {
		if got := obj.Top(); got != want {
			t.Errorf(" obj.Top() = %v, want %v", got, want)
		}
	})

	want = -2
	t.Run("", func(t *testing.T) {
		if got := obj.GetMin(); got != want {
			t.Errorf(" obj.GetMin() = %v, want %v", got, want)
		}
	})

	// ---- case 2
	obj2 := Constructor()

	obj2.Push(-2)
	obj2.Push(0)
	obj2.Push(-1)

	want = -2
	t.Run("", func(t *testing.T) {
		if got := obj2.GetMin(); got != want {
			t.Errorf(" obj2.GetMin() = %v, want %v", got, want)
		}
	})

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj2.Top(); got != want {
			t.Errorf(" obj2.Top() = %v, want %v", got, want)
		}
	})

	obj2.Pop()

	want = -2
	t.Run("", func(t *testing.T) {
		if got := obj2.GetMin(); got != want {
			t.Errorf(" obj2.GetMin() = %v, want %v", got, want)
		}
	})
}
