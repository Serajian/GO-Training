package trains

func IsAnagram(strA, strB string) bool {
	if len(strA) != len(strB) {
		return false
	}
	count := make(map[rune]int)

	for _, char := range strA {
		count[char]++
	}

	for _, char := range strB {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}
	return true

}

//first version:
//func IsAnagram(strA, strB string) bool {
//	if len(strA) != len(strB) {
//		return false
//	}
//	seenA := make(map[rune]int)
//	seenB := make(map[rune]int)
//
//	for _, char := range strA {
//		seenA[char]++
//	}
//	for _, char := range strB {
//		seenB[char]++
//	}
//
//	for _, char := range strA {
//		if seenA[char] != seenB[char] {
//			return false
//		}
//	}
//	return true
//
//}
