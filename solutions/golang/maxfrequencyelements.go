package golang

func maxFrequencyElements(nums []int) int {
	cnt := make(map[int]int)
	maxFreq := 0

	for _, num := range nums {
		cnt[num]++
		maxFreq = Max(maxFreq, cnt[num])
	}

	ans := 0

	for _, freq := range cnt {
		if freq == maxFreq {
			ans += freq
		}
	}

	return ans
}
