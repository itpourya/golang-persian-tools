package commas

func AddCommas(input string) string {
	index := 0
	result := ""

	for _, value := range reverse(input) {
		if index == 3 {
			result += ","
			index = 0
		}
		result += string(value)
		index++
	}
	return reverse(result)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
