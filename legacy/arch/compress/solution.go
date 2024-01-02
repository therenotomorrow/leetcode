package compress

import "strconv"

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	i := 0
	last := chars[0]
	cnt := 1

	for _, char := range chars[1:] {
		if char != last {
			chars[i] = last
			i++

			if cnt > 1 {
				n := strconv.Itoa(cnt)
				for j := range n {
					chars[i] = n[j]
					i++
				}
				cnt = 1
			}

			last = char
		} else {
			cnt++
		}
	}

	chars[i] = last
	i++
	if cnt > 1 {
		n := strconv.Itoa(cnt)

		for j := range n {
			chars[i] = n[j]
			i++
		}
	}

	return i
}
