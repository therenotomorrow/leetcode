package findAllPeople

import (
	"math"

	"github.com/therenotomorrow/leetcode/solutions/golang"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make(map[int][]golang.PairNode)

	for _, meet := range meetings {
		x := meet[0]
		y := meet[1]
		t := meet[2]

		graph[x] = append(graph[x], golang.PairNode{t, y})
		graph[y] = append(graph[y], golang.PairNode{t, x})
	}

	earliest := make([]int, n)
	for i := range earliest {
		earliest[i] = math.MaxInt
	}
	earliest[0] = 0
	earliest[firstPerson] = 0

	q := golang.NewQueue[golang.PairNode]()

	q.Enqueue(golang.PairNode{0, 0})
	q.Enqueue(golang.PairNode{firstPerson, 0})

	// Do BFS
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		person := val[0]
		time := val[1]

		for _, node := range graph[person] {
			t := node[0]
			nextPerson := node[1]

			if t >= time && earliest[nextPerson] > t {
				earliest[nextPerson] = t
				q.Enqueue(golang.PairNode{nextPerson, t})
			}
		}
	}

	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if earliest[i] != math.MaxInt {
			ans = append(ans, i)
		}
	}

	return ans
}
