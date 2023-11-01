package solutions

func decode(encoded []int, first int) []int {
	decoded := make([]int, len(encoded)+1)
	decoded[0] = first

	for i, num := range encoded {
		decoded[i+1] = decoded[i] ^ num
	}

	return decoded
}
