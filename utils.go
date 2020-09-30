package brazil

import "unicode"

func onlyDigits(input string) string {
	norm := make([]rune, 0)

	for _, r := range input {
		if unicode.IsDigit(r) {
			norm = append(norm, r)
		}
	}

	return string(norm)
}

func onlyLetters(input string) string {
	norm := make([]rune, 0)

	for _, r := range input {
		if unicode.IsLetter(r) {
			norm = append(norm, r)
		}
	}

	return string(norm)
}
