package intersection

import "sort"

func intersection(nums [][]int) []int {
	counts := make(map[int]int)
	for _, arr := range nums {
		for _, num := range arr {
			counts[num]++
		}
	}

	ans := make([]int, 0)
	for k, v := range counts {
		if v == len(nums) {
			ans = append(ans, k)
		}
	}

	sort.Ints(ans)

	return ans
}
