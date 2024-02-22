package golang

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := make(map[int]int)

	for _, num := range arr {
		cnt[num]++
	}

	amount := make([]int, 0)
	for _, times := range cnt {
		amount = append(amount, times)
	}

	sort.Ints(amount)

	i := 0

	for i < len(amount) && k > 0 {
		if amount[i] <= k {
			k -= amount[i]
			i++
		} else {
			break
		}
	}

	return len(amount[i:])
}
