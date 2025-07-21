package trains

func ReverseString (str string)string{
	runes:=[]rune(str)
	for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
		runes[i],runes[j]=runes[j],runes[i]
	}
	return string(runes)
}
func ReverseSlice (slice []int)[]int{
	for i,j:=0,len(slice)-1;i<j;i,j=i+1,j-1{
		slice[i],slice[j]=slice[j],slice[i]
	}
	return slice
}