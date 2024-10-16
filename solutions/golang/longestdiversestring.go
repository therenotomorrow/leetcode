package golang

func longestDiverseString(wantA int, wantB int, wantC int) string { //nolint:cyclop
	var (
		currA   = 0
		currB   = 0
		currC   = 0
		runes   = make([]byte, 0)
		zeroing = func(a *int, b *int) { *a = 0; *b = 0 }
		addsub  = func(a *int, b *int) { *a++; *b-- }
	)

	for range wantA + wantB + wantC {
		switch {
		case (wantA >= wantB && wantA >= wantC && currA != 2) || (wantA > 0 && (currB == 2 || currC == 2)):
			runes = append(runes, 'a')

			zeroing(&currB, &currC)
			addsub(&currA, &wantA)
		case (wantB >= wantA && wantB >= wantC && currB != 2) || (wantB > 0 && (currC == 2 || currA == 2)):
			runes = append(runes, 'b')

			zeroing(&currA, &currC)
			addsub(&currB, &wantB)
		case (wantC >= wantA && wantC >= wantB && currC != 2) || (wantC > 0 && (currA == 2 || currB == 2)):
			runes = append(runes, 'c')

			zeroing(&currA, &currB)
			addsub(&currC, &wantC)
		default:
			break
		}
	}

	return string(runes)
}
