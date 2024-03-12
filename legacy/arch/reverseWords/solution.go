package reverseWords

import "strings"

func reverseWords2(s string) string {
	var l, r int

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			l = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			r = i
			break
		}
	}

	sb := strings.Builder{}
	for i, j := r, r; i >= l; {
		for i >= l && s[i] != ' ' {
			i--
		}

		sb.WriteString(s[i+1 : j+1])
		sb.WriteRune(' ')

		for i >= l && s[i] == ' ' {
			i--
		}

		j = i
	}

	return sb.String()[:sb.Len()-1]
}
