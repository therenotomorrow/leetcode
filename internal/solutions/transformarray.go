package solutions

import "slices"

func transformArray(arr []int) []int {
	nums := make([]int, len(arr))
	nums[0] = arr[0]
	nums[len(nums)-1] = arr[len(arr)-1]

	for {
		for i := 1; i < len(arr)-1; i++ {
			nums[i] = arr[i]

			switch {
			case arr[i] < arr[i-1] && arr[i] < arr[i+1]:
				nums[i]++
			case arr[i] > arr[i-1] && arr[i] > arr[i+1]:
				nums[i]--
			}
		}

		if slices.Compare(arr, nums) == 0 {
			break
		}

		copy(arr, nums)
	}

	return arr
}
