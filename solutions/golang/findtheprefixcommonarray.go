package golang

func findThePrefixCommonArray(arr1 []int, arr2 []int) []int {
	const double = 2

	freq := make([]int, len(arr1)+1)
	ans := make([]int, 0, len(arr2))
	cnt := 0

	for i := range arr1 {
		freq[arr1[i]]++
		if freq[arr1[i]] == double {
			cnt++
		}

		freq[arr2[i]]++
		if freq[arr2[i]] == double {
			cnt++
		}

		ans = append(ans, cnt)
	}

	return ans
}
