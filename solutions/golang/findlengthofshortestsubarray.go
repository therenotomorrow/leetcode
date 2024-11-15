package golang

func findLengthOfShortestSubarray(arr []int) int {
	i := 0
	j := len(arr) - 1

	for j > 0 && arr[j-1] <= arr[j] {
		j--
	}

	ans := j

	for i < j && (i == 0 || arr[i-1] <= arr[i]) {
		for j < len(arr) && arr[i] > arr[j] {
			j++
		}

		ans = Min(ans, j-i-1)
		i++
	}

	return ans
}
