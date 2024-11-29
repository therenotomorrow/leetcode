package golang

func canVisitAllRooms(rooms [][]int) bool {
	var (
		ans     = true
		visit   func(key int)
		visited = make([]bool, len(rooms))
	)

	visit = func(key int) {
		if visited[key] || key == len(rooms) {
			return
		}

		visited[key] = true

		for _, room := range rooms[key] {
			visit(room)
		}
	}

	visit(0)

	for _, vis := range visited {
		ans = ans && vis
	}

	return ans
}
