package golang

func findSpecialInteger(arr []int) int {
	window := len(arr) / 4

	for i := 0; i < len(arr)-window; i++ {
		// because array is sorted and exactly one answer
		if arr[i] == arr[i+window] {
			return arr[i]
		}
	}

	return 0
}
