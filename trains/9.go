package trains

import (
	"math"
	"unicode"
)

func MyAToI(s string) int {
	i := 0
	n := len(s)
	sign := 1
	result := 0

	for i < n && s[i] == ' ' {
		i++
	}

	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	for i < n && unicode.IsDigit(rune(s[i])) {
		d := int(s[i] - '0')

		if result > (math.MaxInt32-d)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + d
		i++
	}

	return result * sign
}
