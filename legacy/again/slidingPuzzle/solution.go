package slidingPuzzle

import (
	"strings"
)

func slidingPuzzle(board [][]int) int {
	const target = "123450"

	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	sb := strings.Builder{}
	for _, row := range board {
		for _, val := range row {
			sb.WriteByte(byte(val + '0'))
		}
	}
	start := sb.String()

	queue := []string{start}
	visited := map[string]bool{start: true}
	moves := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == target {
				return moves
			}

			zeroIdx := strings.IndexByte(curr, '0')
			row, col := zeroIdx/3, zeroIdx%3

			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow >= 0 && newRow < 2 && newCol >= 0 && newCol < 3 {
					newIdx := newRow*3 + newCol
					next := swap(curr, zeroIdx, newIdx)
					if !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		moves++
	}

	return -1
}

func swap(s string, i, j int) string {
	runes := []rune(s)
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}
