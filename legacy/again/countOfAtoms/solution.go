package countOfAtoms

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	var (
		index  int
		parser func(formula string) map[string]int
	)

	parser = func(formula string) map[string]int {
		// Local variable
		currMap := make(map[string]int)
		currAtom := ""
		currCount := ""

		// Iterate until the end of the formula
		for index < len(formula) {
			// UPPERCASE LETTER
			if 'A' <= formula[index] && formula[index] <= 'Z' {
				if currAtom != "" {
					if currCount == "" {
						currMap[currAtom]++
					} else {
						val, _ := strconv.Atoi(currCount)
						currMap[currAtom] += val
					}
				}
				currAtom = string(formula[index])
				currCount = ""
				index++
			} else if 'a' <= formula[index] && formula[index] <= 'z' {
				currAtom += string(formula[index])
				index++
			} else if '0' <= formula[index] && formula[index] <= '9' {
				currCount += string(formula[index])
				index++
			} else if formula[index] == '(' {
				index++
				nestedMap := parser(formula)
				for atom, val := range nestedMap {
					currMap[atom] += val
				}
			} else if formula[index] == ')' {
				// Save the last atom and count of nested formula
				if currAtom != "" {
					if currCount == "" {
						currMap[currAtom]++
					} else {
						val, _ := strconv.Atoi(currCount)
						currMap[currAtom] += val
					}
				}

				index++
				multiplier := strings.Builder{}
				for index < len(formula) && '0' <= formula[index] && formula[index] <= '9' {
					multiplier.WriteByte(formula[index])
					index++
				}
				if multiplier.Len() > 0 {
					mult, _ := strconv.Atoi(multiplier.String())
					for atom := range currMap {
						currMap[atom] *= mult
					}
				}

				return currMap
			}
		}

		// Save the last atom and count
		if currAtom != "" {
			if currCount == "" {
				currMap[currAtom]++
			} else {
				val, _ := strconv.Atoi(currCount)
				currMap[currAtom] += val
			}
		}

		return currMap
	}
	// Recursively parse the formula
	finalMap := parser(formula)

	// Sort the final map
	keys := make([]string, 0, len(finalMap))

	for k := range finalMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Generate the answer string
	var ans strings.Builder

	for _, atom := range keys {
		ans.WriteString(atom)
		if finalMap[atom] > 1 {
			ans.WriteString(strconv.Itoa(finalMap[atom]))
		}
	}

	return ans.String()
}
