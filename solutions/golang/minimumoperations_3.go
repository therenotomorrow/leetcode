package golang

func minimumOperations3(nums []int) int {
	const triplet = 3

	ans := 0
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	for ; len(nums) > 0; ans++ {
		distinct := true

		for _, times := range cnt {
			if times != 1 {
				distinct = false

				break
			}
		}

		if distinct {
			break
		}

		size := Min(triplet, len(nums))

		for _, num := range nums[:size] {
			cnt[num]--
			if cnt[num] == 0 {
				delete(cnt, num)
			}
		}

		nums = nums[size:]
	}

	return ans
}
