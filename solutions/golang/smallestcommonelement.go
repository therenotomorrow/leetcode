package golang

func smallestCommonElement(mat [][]int) int {
	search := func(num int, nums []int) bool {
		left := 0
		right := len(nums) - 1

		for left < right {
			mid := (left + right) / Half

			switch {
			case nums[mid] < num:
				left = mid + 1
			case nums[mid] > num:
				right = mid - 1
			default:
				return true
			}
		}

		return num == nums[left]
	}

	for _, num := range mat[0] {
		found := true

		for _, nums := range mat {
			if !search(num, nums) {
				found = false

				break
			}
		}

		if found {
			return num
		}
	}

	return -1
}
