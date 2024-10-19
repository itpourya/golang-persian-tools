package digits

import "strings"

// faToEn converts Persian digits to English digits.
func FaToEn(input string) string {
	return convert(input, 1776, 1728)
}

// enToFa converts English digits to Persian digits.
func EnToFa(input string) string {
	return convert(input, 48, 1728)
}

// enToAr converts English digits to Arabic digits.
func EnToAr(input string) string {
	return convert(input, 48, 1584)
}

// arToEn converts Arabic digits to English digits.
func ArToEn(input string) string {
	return convert(input, 1632, -1584)
}

// faToAr converts Persian digits to Arabic digits.
func FaToAr(input string) string {
	return convert(input, 1776, -144)
}

// arToFa converts Arabic digits to Persian digits.
func ArToFa(input string) string {
	return convert(input, 1632, 144)
}

// convert performs the conversion based on the specified ranges and offset.
func convert(input string, start, offset int) string {
	var result strings.Builder

	for _, char := range input {
		if char >= rune(start) && char < rune(start+10) {
			result.WriteRune(char + rune(offset))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
