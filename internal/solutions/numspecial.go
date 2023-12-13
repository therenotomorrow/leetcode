package solutions

func numSpecial(mat [][]int) int {
	cols := map[int]int{}
	rows := map[int]int{}

	for i, row := range mat {
		for j, cell := range row {
			cols[j] += cell
			rows[i] += cell
		}
	}

	cnt := 0

	for i, row := range mat {
		for j, cell := range row {
			if cell != 1 {
				continue
			}

			if rows[i] == 1 && cols[j] == 1 {
				cnt++
			}
		}
	}

	return cnt
}
