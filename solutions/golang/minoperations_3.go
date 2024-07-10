package golang

func minOperations3(logs []string) int {
	cnt := 0

	for _, log := range logs {
		switch log {
		case "./":
		case "../":
			if cnt > 0 {
				cnt--
			}
		default:
			cnt++
		}
	}

	return cnt
}
