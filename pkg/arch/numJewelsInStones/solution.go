package numJewelsInStones

func numJewelsInStones(jewels string, stones string) int {
	set := make(map[rune]int)
	for _, jewel := range jewels {
		set[jewel] = 1
	}

	var cnt int
	for _, stone := range stones {
		cnt += set[stone]
	}

	return cnt
}
