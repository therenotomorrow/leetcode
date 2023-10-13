package slidingWindow

func findLength(nums []int, k int) int {
	var (
		curr        int // current sum
		res         int // len of longest
		left, right int // indexes of window
	)

	for ; right < len(nums); right++ {
		curr += nums[right]

		for ; curr > k; left++ {
			curr -= nums[left]
		}

		if l := right - left + 1; l > res {
			res = l
		}
	}

	return res
}

func findLengthString(s string) int {
	var (
		curr        int
		res         int
		left, right int
	)

	for ; right < len(s); right++ {
		if s[right] == '0' {
			curr++
		}

		for ; curr > 1; left++ {
			if s[left] == '0' {
				curr--
			}
		}

		if l := right - left + 1; l > res {
			res = l
		}
	}

	return res
}

func findBestSubarray(nums []int, k int) int {
	sum := 0
	for _, num := range nums[:k] {
		sum += num
	}

	res := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > res {
			res = sum
		}
	}

	return res
}
