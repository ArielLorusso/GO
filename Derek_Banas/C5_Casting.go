// 12:12 Casting  https://www.youtube.com/watch?v=YzLrWHZa-Kc

package main

import ("fmt" 	 // format  (Print)
"reflect"
"strconv"
)
var pl = fmt.Println // macro


func main () {
	cV1 := 1.5							// float
	cV2 := int(cV1)
	pl (cV2," -> ",reflect.TypeOf(cV2))	// float cast to string
	
	cV3 := "50000000"					// string
	cV4, err := strconv.Atoi (cV3)
	pl(cV4, err," -> ",reflect.TypeOf(cV4)) // string cast to int
	var cV5 = 0.0
	cV5, err = strconv.ParseFloat(cV3)
	pl(cV5, err," -> ",reflect.TypeOf(cV5)) // string cast to int
}