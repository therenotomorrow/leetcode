package golang

func imageSmoother(img [][]int) [][]int {
	smg := make([][]int, len(img))
	for i, raw := range img {
		smg[i] = make([]int, len(raw))
	}

	for i, row := range img {
		for j := range row {
			var sum, cnt int

			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if 0 <= x && x < len(img) && 0 <= y && y < len(row) {
						sum += img[x][y]
						cnt++
					}
				}
			}

			smg[i][j] = sum / cnt
		}
	}

	return smg
}
