package golang

func isIsomorphic(str string, tpl string) bool {
	mStr := make(map[byte]byte)
	mTpl := make(map[byte]byte)

	for i := range str {
		_, ok1 := mStr[str[i]]
		_, ok2 := mTpl[tpl[i]]

		if !ok1 {
			mStr[str[i]] = tpl[i]
		}

		if !ok2 {
			mTpl[tpl[i]] = str[i]
		}

		if mStr[str[i]] != tpl[i] || mTpl[tpl[i]] != str[i] {
			return false
		}
	}

	return true
}
