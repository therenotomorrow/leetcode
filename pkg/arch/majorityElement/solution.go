package majorityElement

// need use Boyer-Moore Voting Algorithm
func majorityElement(nums []int) []int {
	threshold := len(nums) / 3
	maxNums := 2

	hm := make(map[int]int)
	res := make([]int, 0, maxNums)

	for _, num := range nums {
		if hm[num] == -1 {
			continue
		}

		hm[num]++

		if hm[num] > threshold {
			maxNums--
			res = append(res, num)
			hm[num] = -1
		}

		if maxNums == 0 {
			break
		}
	}

	return res
}
