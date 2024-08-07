package secondMinimum

import "github.com/therenotomorrow/leetcode/solutions/golang"

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	dist1 := make([]int, n+1)
	dist2 := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dist1[i] = -1
		dist2[i] = -1
	}
	queue := golang.NewQueue[[]int]()

	queue.Enqueue([]int{1, 1})
	dist1[1] = 0
	for !queue.IsEmpty() {
		temp, _ := queue.Dequeue()
		node := temp[0]
		freq := temp[1]
		var timeTaken int

		if freq == 1 {
			timeTaken = dist1[node]
		} else {
			timeTaken = dist2[node]
		}

		// If the timeTaken falls under the red bracket, wait till the path turns green.
		if (timeTaken/change)%2 != 0 {
			timeTaken = change*(timeTaken/change+1) + time
		} else {
			timeTaken = timeTaken + time
		}

		for _, neighbor := range adj[node] {
			if dist1[neighbor] == -1 {
				dist1[neighbor] = timeTaken
				queue.Enqueue([]int{neighbor, 1})
			} else if dist2[neighbor] == -1 && dist1[neighbor] != timeTaken {
				if neighbor == n {
					return timeTaken
				}
				dist2[neighbor] = timeTaken
				queue.Enqueue([]int{neighbor, 2})
			}
		}
	}
	return 0
}
