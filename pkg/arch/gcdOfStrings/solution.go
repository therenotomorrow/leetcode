package gcdOfStrings

func gcd(m int, n int) int {
	for m%n != 0 {
		r := m % n
		m, n = n, r
	}

	return n
}

func repeat(s string, p string) bool {
	for i, n := 0, len(p); i < len(s); i += n {
		if s[i:i+n] != p {
			return false
		}
	}

	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)

	if l1 == l2 && str1 != str2 {
		return ""
	}

	pattern := str2[:gcd(l1, l2)]
	if repeat(str1, pattern) && repeat(str2, pattern) {
		return pattern
	}

	return ""
}
