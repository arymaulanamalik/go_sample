package masking

import "unicode"

func MaskingAfterFirstName(str string) string {

	rs := []rune(str)

	var idx int
	for i := 1; i < len(rs); i++ {
		if unicode.IsSpace(rs[i]) {
			idx = i + 1
			if idx <= len(rs)-1 && !unicode.IsSpace(rs[idx]) {
				idx = i + 2
			}
		}
		if idx == i {
			rs[i] = '*'
			idx++
		}
	}

	return string(rs)
}
