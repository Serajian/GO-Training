package trains

import "fmt"

func HowManyAppears(str string) string {
	seen := make(map[rune]int)

	for _, ch := range str {
		seen[ch]++
	}

	var res string
	for key, value := range seen {
		if key == ' ' {
			res += fmt.Sprintf("SPACE: %v\n", value)
		} else {
			res += fmt.Sprintf("%s: %v\n", string(key), value)
		}
	}

	return res
}
