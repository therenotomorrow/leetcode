package golang

type MovingAverage struct {
	queue []int
	size  int
}

func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{queue: make([]int, 0, size), size: size}
}

func (ma *MovingAverage) Next(next int) float64 {
	if len(ma.queue) == ma.size {
		ma.queue = ma.queue[1:]
	}

	ma.queue = append(ma.queue, next)

	sum := 0

	for _, val := range ma.queue {
		sum += val
	}

	return float64(sum) / float64(len(ma.queue))
}
