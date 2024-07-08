package golang

func findTheWinner(n int, k int) int {
	que := NewQueue[int]()
	for i := 1; i <= n; i++ {
		que.Enqueue(i)
	}

	for que.Size() > 1 {
		for step := k - 1; step > 0; step-- {
			val, _ := que.Dequeue()
			que.Enqueue(val)
		}
		que.Dequeue()
	}

	val, _ := que.Peek()

	return val
}
