package trains

func IsPalindrome(str string)bool{
runes:=[]rune(str)
for i,j := 0,len(runes)-1; i<j ; i,j=i+1,j-1 {
	if runes[i]!=runes[j] {
		return false
	}
}
return true
}