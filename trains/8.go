package trains

func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // ذخیره آخرین موقعیت هر حرف
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		// اگر حرف تکراری است و درون پنجره فعلی قرار دارد، left را به جلو ببر
		if lastPos, ok := charIndex[char]; ok && lastPos >= left {
			left = lastPos + 1
		}
		// ذخیره موقعیت آخرین مشاهده حرف
		charIndex[char] = right
		// محاسبه طول پنجره فعلی و مقایسه با maxLen
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
