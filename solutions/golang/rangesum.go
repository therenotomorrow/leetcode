package golang

import "sort"

func rangeSum(nums []int, n int, left int, right int) int {
	newNums := make([]int, 0)

	for i := range n {
		for j := i; j < n; j++ {
			newNums = append(newNums, Sum(nums[i:j+1]...))
		}
	}

	sort.Ints(newNums)

	ans := 0
	for i := left - 1; i < right; i++ {
		ans += newNums[i]
		ans %= MOD
	}

	return ans
}
