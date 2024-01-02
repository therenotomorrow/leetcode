package romanToInt

var mapping = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var num int

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && mapping[s[i]] < mapping[s[i+1]] {
			num += mapping[s[i+1]] - mapping[s[i]]
			i++
		} else {
			num += mapping[s[i]]
		}
	}

	return num
}
