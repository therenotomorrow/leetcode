package solutions

func findArray(pref []int) []int {
	arr := make([]int, len(pref))
	arr[0] = pref[0]

	for i := 1; i < len(pref); i++ {
		// opposite of XOR is XOR
		arr[i] = pref[i-1] ^ pref[i]
	}

	return arr
}
