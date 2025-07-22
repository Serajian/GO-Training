package trains

import "strconv"

func CompressLength(str string) string {
	var compressed string
	count := 1

	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			compressed += string(str[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	compressed += string(str[len(str)-1]) + strconv.Itoa(count)

	return compressed
}
