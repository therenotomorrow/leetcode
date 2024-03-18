package jobScheduling

import (
	"cmp"
	"fmt"
	"sort"
)

func findNextJob(startTime []int, lastEndingTime int) int {
	start := 0
	end := len(startTime) - 1
	nextIndex := len(startTime)

	for start <= end {
		mid := (start + end) / 2
		if startTime[mid] >= lastEndingTime {
			nextIndex = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return nextIndex
}

func findMaxProfit(jobs [][]int, startTime []int, n int, position int, c Cache) int {
	if position == n {
		return 0
	}

	if val, ok := c.Load(position); ok {
		return val
	}

	nextIndex := findNextJob(startTime, jobs[position][1])

	maxProfit := Max(
		findMaxProfit(jobs, startTime, n, position+1, c),
		jobs[position][2]+findMaxProfit(jobs, startTime, n, nextIndex, c),
	)

	c.Save(maxProfit, position)

	return maxProfit
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([][]int, 0)
	c := NewCache()

	// storing job's details into one list
	// this will help in sorting the jobs while maintaining the other parameters
	length := len(profit)
	for i := 0; i < length; i++ {
		currJob := make([]int, 0)
		currJob = append(currJob, startTime[i], endTime[i], profit[i])
		jobs = append(jobs, currJob)
	}

	sort.SliceStable(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	// binary search will be used in startTime so store it as separate list
	for i := 0; i < length; i++ {
		startTime[i] = jobs[i][0]
	}

	return findMaxProfit(jobs, startTime, length, 0, c)
}

func Max[T cmp.Ordered](nums ...T) T {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}

type Cache interface {
	Load(keys ...int) (int, bool)
	Save(val int, keys ...int)
}

type cache struct {
	data map[string]int
}

func NewCache() Cache {
	return &cache{data: make(map[string]int)}
}

func (c *cache) makeKey(parts []int) string {
	return fmt.Sprintf("%v", parts)
}

func (c *cache) Load(keys ...int) (int, bool) {
	val, ok := c.data[c.makeKey(keys)]

	return val, ok
}

func (c *cache) Save(val int, keys ...int) {
	c.data[c.makeKey(keys)] = val
}
