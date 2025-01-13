package golang

func canBeValid(brackets string, locked string) bool {
	const half = 2

	var balance int

	if len(brackets)%half != 0 {
		return false
	}

	balance = 0

	for i, r := range brackets {
		if locked[i] == '0' || r == '(' {
			balance++
		} else {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	balance = 0

	for i := len(brackets) - 1; i >= 0; i-- {
		if locked[i] == '0' || brackets[i] == ')' {
			balance++
		} else {
			balance--
		}

		if balance < 0 {
			return false
		}
	}

	return true
}
