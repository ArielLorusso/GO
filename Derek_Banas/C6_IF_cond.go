// 17:33 If Conditional  https://www.youtube.com/watch?v=YzLrWHZa-Kc

package main

import ("fmt" 	 // format  (Print)

)
var pl = fmt.Println // macro


func main () {
	// Conditional Operators : > < >= <= == !=
	// Logical Operators 	 :  && 	|| 	!
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday") 
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not and Important Birthday")
	}
}