package trains

import "strings"

func FlipTheWords(text string) string {
	arr := strings.Fields(text)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, " ")
}
