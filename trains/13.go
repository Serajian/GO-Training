package trains

import "strings"

func IsRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	repeatStr := str1 + str1

	result := strings.Contains(repeatStr, str2)

	return result
}
