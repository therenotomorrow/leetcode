package golang

import "sort"

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	ans := make([]int, 0)
	order := make([]int, 0)
	maxes := make(map[int][]int)

	for num, times := range cnt {
		maxes[times] = append(maxes[times], num)
		order = append(order, times)
	}

	sort.Stable(sort.Reverse(sort.IntSlice(order)))

	for i := 0; i < k; {
		for _, num := range maxes[order[i]] {
			ans = append(ans, num)
			i++
		}
	}

	return ans
}
