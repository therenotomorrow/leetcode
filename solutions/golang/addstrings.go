package golang

import "slices"

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}

	var (
		i        = len(num1) - 1
		j        = len(num2) - 1
		bytes    = make([]byte, 0)
		reminder int
	)

	for j >= 0 {
		digit1 := int(num1[i] - '0')
		digit2 := int(num2[j] - '0')

		value := (digit1 + digit2 + reminder) % Digits
		reminder = (digit1 + digit2 + reminder) / Digits

		bytes = append(bytes, byte(value)+'0')

		i--
		j--
	}

	for i >= 0 {
		digit1 := int(num1[i] - '0')

		value := (digit1 + reminder) % Digits
		reminder = (digit1 + reminder) / Digits

		bytes = append(bytes, byte(value)+'0')

		i--
	}

	if reminder > 0 {
		bytes = append(bytes, byte(reminder)+'0')
	}

	slices.Reverse(bytes)

	return string(bytes)
}
