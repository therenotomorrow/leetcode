package firstBadVersion

func isBadVersion(version int) bool {
	switch version {
	case 1, 4, 5:
		return true
	default:
		return false
	}
}

func firstBadVersion(n int) int {
	a, b := 1, n

	for a < b {
		if ver := a + (b-a)/2; isBadVersion(ver) {
			b = ver
		} else {
			a = ver + 1
		}
	}

	return a
}
