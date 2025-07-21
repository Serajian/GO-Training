package main

import (
	"fmt"
	"go-train/trains"
)

func main(){
	fmt.Println("--------HI from MAIN---------")

	tets:="Hello, World!"
	// tets:="stress"
	res:=trains.RemoveDuplicates(tets)
	fmt.Println("RESULT: ",res)
	
	fmt.Println("********BYE*********")
}