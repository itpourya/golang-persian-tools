package extractcardnumber

import (
	"strings"
)

const (
	enNumber = "0123456789"
)

func ExtractCardNumber(input string) string {
	var card string

	for _, char := range input {
		if strings.Contains(enNumber, string(char)) {
			card += string(char)
		}
		if string(char) == "-" {
			card += string(char)
		}
	}

	return card
}
