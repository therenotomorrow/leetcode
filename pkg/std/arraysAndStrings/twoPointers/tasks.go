package twoPointers

func checkIfPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func checkForTarget(nums []int, target int) bool {
	i := 0
	j := len(nums) - 1

	for i < j {
		sum := nums[i] + nums[j]

		switch {
		case sum > target:
			j--
		case sum < target:
			i++
		default:
			return true
		}
	}

	return false
}

func combine(arr1 []int, arr2 []int) []int {
	var (
		i, j, k = 0, 0, 0
		res     = make([]int, len(arr1)+len(arr2))
	)

	for ; i < len(arr1) && j < len(arr2); k++ {
		if arr1[i] < arr2[j] {
			res[k] = arr1[i]
			i++
		} else {
			res[k] = arr2[j]
			j++
		}
	}

	for ; i < len(arr1); i++ {
		res[k] = arr1[i]
		k++
	}

	for ; j < len(arr2); j++ {
		res[k] = arr2[j]
		k++
	}

	return res
}
