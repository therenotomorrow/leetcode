package golang

import "strings"

func customSortString(order string, s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	sb := strings.Builder{}

	for _, r := range order {
		cnt := m[r]

		for i := 0; i < cnt; i++ {
			sb.WriteRune(r)
		}

		delete(m, r)
	}

	for r, cnt := range m {
		for i := 0; i < cnt; i++ {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
