package golang

func getFood(grid [][]byte) int { //nolint:cyclop
	var (
		start   PairNode
		canMove = func(point PairNode) bool {
			return 0 <= point[0] && point[0] < len(grid) &&
				0 <= point[1] && point[1] < len(grid[0]) &&
				grid[point[0]][point[1]] != 'X'
		}
		directions = []PairNode{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	)

	for i, row := range grid {
		for j, cell := range row {
			if cell == '*' {
				// will be the only point
				start[0] = i
				start[1] = j
			}
		}
	}

	steps := 1
	que := NewQueue[PairNode]()

	que.Enqueue(start)

	for !que.IsEmpty() {
		size := que.Size()

		for range size {
			point, _ := que.Dequeue()

			for _, direction := range directions {
				newPoint := PairNode{point[0] + direction[0], point[1] + direction[1]}

				if !canMove(newPoint) {
					continue
				}

				if grid[newPoint[0]][newPoint[1]] == '#' {
					return steps
				}

				grid[newPoint[0]][newPoint[1]] = 'X'

				que.Enqueue(newPoint)
			}
		}

		steps++
	}

	return -1
}
