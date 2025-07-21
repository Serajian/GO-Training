package trains

import "strings"

func FirstNoneRepeating(str string)string{
	str=strings.ToLower(str)
	seen:= make(map[rune]int)

	for _,char:=range str{
		seen[char]++
	}

	for _,char:=range str{
		count:=seen[char]
		if count==1 {
			return string(char)
		}
	}
	return "NO REPEATING"
}