package golang

import "sort"

func leastInterval(tasks []byte, n int) int {
	// heap could be here
	freq := make([]int, Alphabet)
	for _, task := range tasks {
		freq[task-'A']++
	}

	sort.Ints(freq)

	maxF := freq[Alphabet-1] - 1 // intervals
	idle := maxF * n             // max of idle times could be

	for i := 24; i >= 0 && freq[i] > 0; i-- {
		idle -= Min(maxF, freq[i]) // next of idle could be after max
	}

	if idle > 0 {
		return idle + len(tasks)
	}

	return len(tasks)
}
