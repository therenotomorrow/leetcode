package golang

func countSubIslands(grid1 [][]int, grid2 [][]int) int { //nolint:cyclop
	var (
		cnt     = 0
		rows    = len(grid2)
		cols    = len(grid2[0])
		dirs    = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		visited = make([][]bool, rows)
		bfs     = func(startRow int, startCol int) bool {
			found := true
			que := NewQueue[PairNode]()

			visited[startRow][startCol] = true

			for que.Enqueue(PairNode{startRow, startCol}); !que.IsEmpty(); {
				node, _ := que.Dequeue()

				row, col := node[0], node[1]

				if grid1[row][col] != 1 {
					// we should travers the whole island for decide
					found = false
				}

				for _, dir := range dirs {
					newRow := row + dir[0]
					newCol := col + dir[1]

					if newRow >= 0 &&
						newCol >= 0 &&
						newRow < rows &&
						newCol < cols &&
						grid2[newRow][newCol] == 1 &&
						!visited[newRow][newCol] {
						que.Enqueue(PairNode{newRow, newCol})

						visited[newRow][newCol] = true
					}
				}
			}

			return found
		}
	)

	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for row := range grid2 {
		for col := range grid2[row] {
			if !visited[row][col] && grid2[row][col] == 1 && bfs(row, col) {
				cnt++
			}
		}
	}

	return cnt
}
