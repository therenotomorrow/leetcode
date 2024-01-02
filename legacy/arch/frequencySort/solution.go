package frequencySort

import "strings"

func frequencySort(s string) string {
	max := 0
	fm := make(map[rune]int)
	for _, r := range s {
		fm[r]++

		if val := fm[r]; val > max {
			max = val
		}
	}

	order := make([][]string, max+1)
	for r, c := range fm {
		order[c] = append(order[c], strings.Repeat(string(r), c))
	}

	sb := strings.Builder{}
	for i := max; i >= 0; i-- {
		for _, s := range order[i] {
			sb.WriteString(s)
		}
	}

	return sb.String()
}
