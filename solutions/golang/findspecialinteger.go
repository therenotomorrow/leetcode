package golang

func findSpecialInteger(arr []int) int {
	const size = 4

	window := len(arr) / size

	for i := range len(arr) - window {
		// because array is sorted and exactly one answer
		if arr[i] == arr[i+window] {
			return arr[i]
		}
	}

	return 0
}
