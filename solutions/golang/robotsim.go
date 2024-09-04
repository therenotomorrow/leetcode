package golang

func robotSim(commands []int, obstacles [][]int) int {
	const (
		turnRight  = -1
		turnLeft   = -2
		multiplier = 1_000_000
	)

	var (
		dir     = 0
		dirs    = [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
		point   = []int{0, 0}
		maxDist = 0
	)

	obstaclesSet := NewSet[int]()
	for _, obstacle := range obstacles {
		obstaclesSet.Add(obstacle[0] + multiplier*obstacle[1])
	}

	for _, command := range commands {
		switch command {
		case turnRight:
			dir--
			if dir < 0 {
				dir = len(dirs) - 1
			}
		case turnLeft:
			dir++
			if dir >= len(dirs) {
				dir = 0
			}
		default:
			for range command {
				newX, newY := point[0]+dirs[dir][0], point[1]+dirs[dir][1]

				if obstaclesSet.Contains(newX + multiplier*newY) {
					break
				}

				point[0], point[1] = newX, newY
			}
		}

		maxDist = Max(maxDist, point[0]*point[0]+point[1]*point[1])
	}

	return maxDist
}
