package main

import (
	"fmt" 	 // format  (Print)
)

var pl = fmt.Println // macro


func main () {
	var vName string = "Derek"
	var v1 ,v2 = 1.2,3.4
	// v3  = "hello"		<-- ERROR inicialization without declaration
	// v3 := "hello" 		<-- OK    declaration + asignment
	// v3 string := "hello" <-- ERROR := cant have type
	var v3 = "hello"
	v3 = "godbye"
	pl(vName,v1,v3,v2)
}