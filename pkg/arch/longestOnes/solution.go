package longestOnes

func longestOnes(nums []int, k int) int {
	curr := 0
	res := 0

	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			curr += 1
		}

		for ; curr > k; l++ {
			if nums[l] == 0 {
				curr -= 1
			}
		}

		if l := r - l + 1; l > res {
			res = l
		}
	}

	return res
}
