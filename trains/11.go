package trains

func MostFrequentCharacter(str string) string {
	seen := make(map[rune]int)
	maxCount := 0
	var result string

	for _, char := range str {
		seen[char]++
		if seen[char] > maxCount {
			maxCount = seen[char]
			result = string(char)
		}
	}
	return result
}
