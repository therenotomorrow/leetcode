package Graph

import "testing"

func TestSolution(t *testing.T) {
	var obj Graph
	var want int

	obj = Constructor(4, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}})

	want = 6
	t.Run("", func(t *testing.T) {
		if got := obj.ShortestPath(3, 2); got != want {
			t.Errorf("obj.ShortestPath() = %v, want = %v", got, want)
		}
	})

	want = -1
	t.Run("", func(t *testing.T) {
		if got := obj.ShortestPath(0, 3); got != want {
			t.Errorf("obj.ShortestPath() = %v, want = %v", got, want)
		}
	})

	obj.AddEdge([]int{1, 3, 4})

	want = 6
	t.Run("", func(t *testing.T) {
		if got := obj.ShortestPath(0, 3); got != want {
			t.Errorf("obj.ShortestPath() = %v, want = %v", got, want)
		}
	})
}
