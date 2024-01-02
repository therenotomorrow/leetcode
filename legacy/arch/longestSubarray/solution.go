package longestSubarray

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubarray(nums []int) int {
	var cnt, res, i, j int

	for j < len(nums) {
		if nums[j] == 0 {
			cnt++
		}

		for cnt > 1 {
			if nums[i] == 0 {
				cnt--
			}
			i++
		}

		res = max(j-i, res)
		j++
	}

	if res == len(nums) {
		return res - 1
	}

	return res
}

func longestSubarray2(nums []int, limit int) int {
	var (
		inc  = make([]int, 0)
		dec  = make([]int, 0)
		left = 0
		ans  = 0
	)

	for right, num := range nums {
		for l := len(inc); l > 0 && inc[l-1] > num; l-- {
			inc = inc[:l-1]
		}
		for l := len(dec); l > 0 && dec[l-1] < num; l-- {
			dec = dec[:l-1]
		}

		inc = append(inc, num)
		dec = append(dec, num)

		for dec[0]-inc[0] > limit {
			if nums[left] == dec[0] {
				dec = dec[1:]
			}
			if nums[left] == inc[0] {
				inc = inc[1:]
			}
			left += 1
		}

		if tmp := right - left + 1; tmp > ans {
			ans = tmp
		}
	}

	return ans
}
