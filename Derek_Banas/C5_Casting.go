// 12:12 Casting  https://www.youtube.com/watch?v=YzLrWHZa-Kc

package main

import ("fmt" 	 // format  (Print)
"reflect"
"strconv"
)
var pl = fmt.Println // macro


func main () {
	cV1 := 1.5							// float
	cV2 := int(cV1)							// float cast to string
	pl (cV2," -> ",reflect.TypeOf(cV2))	
	
	cV3 := "50000000"					// string
	cV4, err := strconv.Atoi (cV3)
	pl("err is -> ",reflect.TypeOf(err)) 	// interface {Error() string}   : <nil>
	pl(cV4," -> ",reflect.TypeOf(cV4)) 
	cV5, _ := strconv.ParseFloat(cV3,64)	// string cast to float
	pl(cV5," -> ",reflect.TypeOf(cV5)) 
	
	cV6 := fmt.Sprintf("%f %d",3.14,10)	    // format  float & int to String
	pl(cV6," -> ",reflect.TypeOf(cV6)) 
}