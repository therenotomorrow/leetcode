package golang

import "strconv"

func validWordAbbreviation(word string, abbr string) bool {
	var i, digit int

	for j := range len(abbr) {
		val, err := strconv.Atoi(string(abbr[j]))

		if err == nil {
			digit *= Digits
			digit += val

			if digit == 0 {
				return false // leading zeros
			}

			continue
		}

		i += digit
		digit = 0

		if i >= len(word) || word[i] != abbr[j] {
			return false
		}

		i++
	}

	if digit != 0 {
		i += digit
	}

	return i == len(word)
}
