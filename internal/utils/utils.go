package utils

import "strconv"

// formatWithCommas adds thousand separators to an integer.
func FormatWithCommas(n int) string {
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}
	var result string
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(r)
	}
	return result
}
