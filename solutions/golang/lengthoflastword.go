package golang

func lengthOfLastWord(str string) int {
	cnt := 0
	right := len(str) - 1

	for str[right] == ' ' {
		right--
	}

	for ; right >= 0 && str[right] != ' '; right-- {
		cnt++
	}

	return cnt
}
