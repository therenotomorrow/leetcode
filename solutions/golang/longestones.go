package golang

func longestOnes(nums []int, k int) int {
	var curr, ans, left int

	for right, num := range nums {
		if num == 0 {
			curr++
		}

		for ; curr > k; left++ {
			if nums[left] == 0 {
				curr--
			}
		}

		ans = Max(ans, right-left+1)
	}

	return ans
}
