package predictPartyVictory

func predictPartyVictory(senate string) string {
	radiant := make([]int, 0)
	dire := make([]int, 0)

	for i, sen := range senate {
		if sen == 'D' {
			dire = append(dire, i)
		} else {
			radiant = append(radiant, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		d := dire[0]
		r := radiant[0]

		dire = dire[1:]
		radiant = radiant[1:]

		if n := len(senate); d < r {
			dire = append(dire, d+n)
		} else {
			radiant = append(radiant, r+n)
		}
	}

	if len(radiant) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}
