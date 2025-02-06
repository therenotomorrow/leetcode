package golang

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	cnt := 0

	for _, item := range items {
		isMatch := false

		switch ruleKey {
		case "type":
			isMatch = item[0] == ruleValue
		case "color":
			isMatch = item[1] == ruleValue
		case "name":
			isMatch = item[2] == ruleValue
		}

		if isMatch {
			cnt++
		}
	}

	return cnt
}
