package golang

func minimumRecolors(blocks string, k int) int {
	var ans, cnt, right int

	for ; right < k; right++ {
		if blocks[right] == 'W' {
			cnt++
		}
	}

	ans = cnt

	for left := 0; right < len(blocks); {
		if blocks[left] == 'W' {
			cnt--
		}

		if blocks[right] == 'W' {
			cnt++
		}

		ans = Min(ans, cnt)

		left++
		right++
	}

	return ans
}
