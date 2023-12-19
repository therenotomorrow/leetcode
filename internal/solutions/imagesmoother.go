package solutions

func imageSmoother(img [][]int) [][]int {
	smg := make([][]int, len(img))
	for i, raw := range img {
		smg[i] = make([]int, len(raw))
	}

	for i, row := range img {
		for j := range row {
			sum := 0
			cnt := 0

			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if 0 <= x && x < len(img) && 0 <= y && y < len(row) {
						sum += img[x][y]
						cnt += 1
					}
				}
			}

			smg[i][j] = sum / cnt
		}
	}

	return smg
}
