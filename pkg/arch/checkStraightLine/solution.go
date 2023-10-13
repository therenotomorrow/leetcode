package checkStraightLine

func checkStraightLine(coordinates [][]int) bool {
	diffX := coordinates[1][0] - coordinates[0][0]
	diffY := coordinates[1][1] - coordinates[0][1]

	for i := 1; i < len(coordinates); i++ {
		if diffY*(coordinates[i][0]-coordinates[0][0]) != diffX*(coordinates[i][1]-coordinates[0][1]) {
			return false
		}
	}

	return true
}
