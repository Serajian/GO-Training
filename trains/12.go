package trains

func MakeSubstrings(str string) []string {
	var substrings []string
	n := len(str)

	for start := 0; start < n; start++ {
		for end := start + 1; end <= n; end++ {
			sub := str[start:end]
			substrings = append(substrings, sub)
		}
	}
	return substrings
}
