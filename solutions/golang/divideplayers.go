package golang

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	canonSum := skill[0] + skill[len(skill)-1]
	chemistry := int64(0)

	for i, j := 0, len(skill)-1; i < j; {
		if skill[i]+skill[j] != canonSum {
			return -1
		}

		chemistry += int64(skill[i] * skill[j])

		i++
		j--
	}

	return chemistry
}
