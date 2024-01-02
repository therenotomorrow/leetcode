package addBinary

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return add(a, b)
	}
	return add(b, a)
}

func add(a string, b string) string {
	var (
		bytes  []byte
		remain bool
	)

	pa := len(a) - 1
	pb := len(b) - 1

	for ; pa >= 0; pa, pb = pa-1, pb-1 {
		l, r := a[pa], b[pb]

		switch {
		case l == '1' && r == '1':
			if remain {
				bytes = append(bytes, '1')
			} else {
				bytes = append(bytes, '0')
			}
			remain = true
		case l == '1' || r == '1':
			if remain {
				bytes = append(bytes, '0')
			} else {
				bytes = append(bytes, '1')
				remain = false
			}
		default:
			if remain {
				bytes = append(bytes, '1')
			} else {
				bytes = append(bytes, '0')
			}
			remain = false
		}
	}

	for ; pb >= 0; pb-- {
		r := b[pb]

		switch {
		case r == '1':
			if remain {
				bytes = append(bytes, '0')
			} else {
				bytes = append(bytes, '1')
				remain = false
			}
		default:
			if remain {
				bytes = append(bytes, '1')
			} else {
				bytes = append(bytes, '0')
			}
			remain = false
		}
	}

	if remain {
		bytes = append(bytes, '1')
	}

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}

	return string(bytes)
}
