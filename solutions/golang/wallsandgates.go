package golang

import "math"

func wallsAndGates(rooms [][]int) { //nolint:cyclop
	const (
		gate  = 0
		empty = math.MaxInt32
	)

	var (
		que     = NewQueue[[]int]()
		dirs    = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		allowed = func(row int, col int) bool {
			return 0 <= row && row < len(rooms) && 0 <= col && col < len(rooms[0]) && rooms[row][col] == empty
		}
	)

	for row, line := range rooms {
		for col, room := range line {
			if room != gate {
				continue
			}

			que.Enqueue([]int{row, col})
		}
	}

	for !que.IsEmpty() {
		coords, _ := que.Dequeue()
		row, col := coords[0], coords[1]

		for _, dir := range dirs {
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if allowed(nextRow, nextCol) {
				rooms[nextRow][nextCol] = rooms[row][col] + 1

				que.Enqueue([]int{nextRow, nextCol})
			}
		}
	}
}
