package golang

func sortArray(nums []int) []int {
	const half = 2

	var (
		sort  func(arr []int) []int
		merge func(left []int, right []int) []int
	)

	sort = func(arr []int) []int {
		if len(arr) < half {
			return arr
		}

		if len(arr) == half {
			if arr[0] > arr[1] {
				return []int{arr[1], arr[0]}
			}

			return []int{arr[0], arr[1]}
		}

		mid := len(arr) / half

		return merge(sort(arr[:mid]), sort(arr[mid:]))
	}

	merge = func(left []int, right []int) []int {
		i := 0
		j := 0
		k := 0
		arr := make([]int, len(left)+len(right))

		for ; i < len(left) && j < len(right); k++ {
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
		}

		for ; i < len(left); i++ {
			arr[k] = left[i]
			k++
		}

		for ; j < len(right); j++ {
			arr[k] = right[j]
			k++
		}

		return arr
	}

	return sort(nums)
}
