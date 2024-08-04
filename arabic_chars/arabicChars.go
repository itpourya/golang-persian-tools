package arabic_chars

import (
	"strings"
)

var (
	arNumbers = "٠١٢٣٤٥٦٧٨٩"
	arChars   = "ابتثجحخدذرزسشصضطظعغفقكلمنهويآةى؟ؠءأؤإ ؘ ؙ ؚ؛ ً ٌ ٍ َ ُ ِ ّ ْ ٓ ٔ ٕ ٖ ٗ ٘ ٙ ٚ ٛ ٝ ٞ ٟ٠١٢٣٤٥٦٧٨٩"
)

func HasArabicChars(input string) bool {
	for _, value := range arChars {
		if strings.Contains(input, string(value)) {
			return true
		}
	}

	return false
}

func HasArabicNumbers(input string) bool {
	for _, value := range arNumbers {
		if strings.Contains(input, string(value)) {
			return true
		}
	}

	return false
}

func ToArabicChars(input string) string {
	newText := strings.ReplaceAll(input, string('ك'), "ك")
	newText = strings.ReplaceAll(newText, "ي", "ی")
	newText = strings.ReplaceAll(newText, "ى", "ی")

	return newText
}
