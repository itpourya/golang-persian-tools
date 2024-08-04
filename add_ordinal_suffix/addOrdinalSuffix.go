package add_ordinal_suffix

import (
	"strings"
)

func AddOrdinalSuffix(input string) string {
	if strings.HasSuffix(input, "ی") {
		return input + "اُم"
	} else if strings.HasSuffix(input, "سه") {
		return input[0:len(input)-2] + "وم"
	}
	return input + "ُم"
}
