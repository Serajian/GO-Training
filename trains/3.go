package trains

func RemoveDuplicates(text string) string {
    seen := make(map[rune]bool)
    var result []rune
    
    for _, char := range text {
		_,ok:=seen[char]

		if !ok{
    	seen[char] = true
      result = append(result, char)
		}
    }
    
    return string(result)
}

