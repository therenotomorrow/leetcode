package golang

func countStudents(students []int, sandwiches []int) int {
	var circulars, squares int

	for _, student := range students {
		if student == 0 {
			circulars++
		} else {
			squares++
		}
	}

	for _, sandwich := range sandwiches {
		if sandwich == 0 && circulars == 0 {
			return squares
		}

		if sandwich == 1 && squares == 0 {
			return circulars
		}

		if sandwich == 0 {
			circulars--
		} else {
			squares--
		}
	}

	return 0
}
