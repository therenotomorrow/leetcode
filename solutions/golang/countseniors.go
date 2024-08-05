package golang

import "strconv"

func countSeniors(details []string) int {
	// by task description
	const (
		ageLim = 60
		ageIdx = 11
	)

	numOfPassengers := 0

	for _, detail := range details {
		age, _ := strconv.Atoi(detail[ageIdx : ageIdx+2])
		if age > ageLim {
			numOfPassengers++
		}
	}

	return numOfPassengers
}
