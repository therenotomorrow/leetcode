package golang

func minCost2(costs [][]int) int {
	const (
		color0 = iota
		color1
		color2
	)

	var (
		cache   = NewCache()
		dynamic func(color int, house int) int
	)

	dynamic = func(color int, house int) int {
		if house >= len(costs) {
			return 0
		}

		if val, ok := cache.Load(color, house); ok {
			return val
		}

		cost := costs[house][color]

		switch color {
		case color0:
			cost += Min(dynamic(color1, house+1), dynamic(color2, house+1))
		case color1:
			cost += Min(dynamic(color0, house+1), dynamic(color2, house+1))
		case color2:
			cost += Min(dynamic(color0, house+1), dynamic(color1, house+1))
		}

		cache.Save(cost, color, house)

		return cost
	}

	return Min(dynamic(color0, 0), dynamic(color1, 0), dynamic(color2, 0))
}
