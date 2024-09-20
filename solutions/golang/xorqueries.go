package golang

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0

	for i := range arr {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

	ret := make([]int, len(queries))

	for i, query := range queries {
		left := query[0]
		right := query[1] + 1 // use as range (before was +1 to `prefix`)

		ret[i] = prefix[right] ^ prefix[left]
	}

	return ret
}
